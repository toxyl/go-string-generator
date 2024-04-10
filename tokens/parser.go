package tokens

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/toxyl/go-string-generator/utils"
)

var reRandHash *regexp.Regexp = regexp.MustCompile(`#\d+`)
var reRandInt *regexp.Regexp = regexp.MustCompile(`\[\d+-\d+\]`)
var reIntRange *regexp.Regexp = regexp.MustCompile(`\[\d+\.\.\d+\]`)
var reRandStr *regexp.Regexp = regexp.MustCompile(`(.+?,)+([^,]+)`)

const (
	TOKEN_START              = "["
	TOKEN_END                = "]"
	TOKEN_TYPE_STR           = "str"
	TOKEN_TYPE_STR_L         = TOKEN_TYPE_STR + ":"
	TOKEN_TYPE_STR_U         = TOKEN_TYPE_STR + "U:"
	TOKEN_TYPE_STR_R         = TOKEN_TYPE_STR + "R:"
	TOKEN_TYPE_ALPHA         = "mix"
	TOKEN_TYPE_ALPHA_L       = TOKEN_TYPE_ALPHA + ":"
	TOKEN_TYPE_ALPHA_U       = TOKEN_TYPE_ALPHA + "U:"
	TOKEN_TYPE_ALPHA_R       = TOKEN_TYPE_ALPHA + "R:"
	TOKEN_TYPE_INT           = "int:"
	TOKEN_TYPE_UUID          = "#UUID"
	TOKEN_TYPE_BASE64_ENCODE = "b64:"
	TOKEN_TYPE_URL_ENCODE    = "url:"
	TOKEN_TYPE_FROM_FILE     = ":"
)

type Token string
type UnknownToken string

type RandomStringGenerator struct {
	DataDir string
}

func (rsg *RandomStringGenerator) token(data []byte, eof bool) (int, []byte, error) {
	inToken := false
	returnToken := false
	returnTokenData := false
	offset := 0
	index := 0
	depth := 0
	var char byte
	tokenOpen := TOKEN_START[0]
	tokenClose := TOKEN_END[0]
	for index, char = range data {
		if returnTokenData {
			return index, data[offset:index], nil // return the token data
		}

		if char == tokenOpen {
			if !inToken {
				inToken = true
				offset = index
				if index > 0 {
					returnToken = true
				}
			}
			depth++
		} else if char == tokenClose {
			if inToken && depth == 1 {
				returnTokenData = true
			}
			depth--
		}

		if returnToken {
			return index, data[:index], nil // return the token
		}
	}

	if index > 0 {
		if inToken {
			v := data[:index+1]
			return index, v, nil // return the token
		}

		return index, data[:index], nil // we get here if the string does not end in a token
	}

	return 0, nil, nil // there was no data
}

func (rsg *RandomStringGenerator) tokenize(pattern string) []interface{} {
	pattern = utils.RemoveNonPrintable(pattern)
	if pattern[len(pattern)-1:] != " " {
		pattern += " " // fixes a case where the end of the string is not correctly returned
	}
	s := bufio.NewScanner(strings.NewReader(pattern))

	s.Split(rsg.token)
	tokens := []interface{}{}
	for s.Scan() {
		t := s.Text()
		l := len(t) - 1
		iv := t[:1] == TOKEN_START && t[l:] == TOKEN_END && strings.Count(t, TOKEN_START) == 1
		if iv {
			if l >= 4 && t[1:4] == TOKEN_TYPE_STR {
				if t[1:5] == TOKEN_TYPE_STR_L {
					tokens = append(tokens, TokenStrLower{Length: utils.IntFromString(t[5:l], 0)})
					continue
				}

				if t[1:6] == TOKEN_TYPE_STR_U {
					tokens = append(tokens, TokenStrUpper{Length: utils.IntFromString(t[6:l], 0)})
					continue
				}

				if t[1:6] == TOKEN_TYPE_STR_R {
					tokens = append(tokens, TokenStr{Length: utils.IntFromString(t[6:l], 0)})
					continue
				}
			}

			if l >= 4 && t[1:4] == TOKEN_TYPE_ALPHA {
				if t[1:5] == TOKEN_TYPE_ALPHA_L {
					tokens = append(tokens, TokenMixLower{Length: utils.IntFromString(t[5:l], 0)})
					continue
				}

				if t[1:6] == TOKEN_TYPE_ALPHA_U {
					tokens = append(tokens, TokenMixUpper{Length: utils.IntFromString(t[6:l], 0)})
					continue
				}

				if t[1:6] == TOKEN_TYPE_ALPHA_R {
					tokens = append(tokens, TokenMix{Length: utils.IntFromString(t[6:l], 0)})
					continue
				}
			}

			if l >= 2 && t[1:2] == "#" {
				if l >= 6 && t[1:6] == TOKEN_TYPE_UUID {
					tokens = append(tokens, TokenRandomUUID(""))
					continue
				}

				if reRandHash.MatchString(t) {
					tokens = append(tokens, TokenHash{Length: utils.IntFromString(t[2:l], 0)})
					continue
				}
			}

			if l >= 6 && t[1:5] == TOKEN_TYPE_INT {
				tokens = append(tokens, TokenInt{Length: utils.IntFromString(t[5:l], 0)})
				continue
			}

			if strings.ContainsAny(t, "1234567890.-") {
				if reRandInt.MatchString(t) {
					v := strings.Split(t[1:l], "-")
					tokens = append(tokens, TokenIntRange{Min: utils.IntFromString(v[0], 0), Max: utils.IntFromString(v[1], 0)})
					continue
				}

				if reIntRange.MatchString(t) {
					v := strings.Split(t[1:l], "..")
					tokens = append(tokens, TokenIntList{Min: utils.IntFromString(v[0], 0), Max: utils.IntFromString(v[1], 0)})
					continue
				}
			}

			if t[1:2] == TOKEN_TYPE_FROM_FILE {
				tokens = append(tokens, TokenLineFromFile{File: t[2:l]})
				continue
			}

			if l >= 6 {
				if t[1:5] == TOKEN_TYPE_BASE64_ENCODE {
					b64 := base64.URLEncoding.EncodeToString([]byte(t[5:l]))
					tokens = append(tokens, Token(b64))
					continue
				}

				if t[1:5] == TOKEN_TYPE_URL_ENCODE {
					tokens = append(tokens, Token(url.QueryEscape(t[5:l])))
					continue
				}
			}

			if reRandStr.MatchString(t) {
				tokens = append(tokens, TokenStrFromList{Strings: strings.Split(t[1:l], ",")})
				continue
			}

			tokens = append(tokens, UnknownToken(t))
			continue
		}
		tokens = append(tokens, Token(t))
	}
	return tokens
}

// Generate takes the input string and replaces all tokens
// with randomly generated data. Tokens can be nested.
//
// Available tokens:
//
//	[#UUID]   = random UUID (xxxxxxxx-xxxx-xxxx-xxxxxxxxxxxx)
//	[#56]     = random 56-characters hash
//	[int:6]   = random 6-characters integer (zero-padded)
//	[str:6]   = random 6-characters lowercase string (a-z)
//	[strU:6]  = random 6-characters uppercase string (A-Z)
//	[strR:6]  = random 6-characters mixed-case string (a-z, A-Z)
//	[mix:6]   = random 6-characters lowercase alphanumeric string (a-z, 0-9)
//	[mixU:6]  = random 6-characters uppercase alphanumeric string (A-Z, 0-9)
//	[mixR:6]  = random 6-characters mixed-case alphanumeric string (a-z, A-Z, 0-9)
//	[10-500]  = random value between 10 and 500 (inclusive)
//	[10..500] = comma separated list with all ints from 10 to 500 (inclusive)
//	[a,b,c]   = random value from the list
//
//	[b64:data] = base64-encodes 'data'
//	[url:data] = url-encodes 'data'
//	[:path]    = reads a random line from the given "path" (if path is a directory a random file from that directory will be used)
func (pp *RandomStringGenerator) Generate(str string) string {
	if str == "" {
		return str
	}

	tokens := pp.tokenize(str)
	parsed := []string{}
	for _, t := range tokens {
		switch tt := t.(type) {
		case TokenLineFromFile:
			parsed = append(parsed, pp.Generate(tt.Parse(pp.DataDir)))
		case TokenInt:
			parsed = append(parsed, tt.Parse())
		case TokenIntRange:
			parsed = append(parsed, tt.Parse())
		case TokenIntList:
			parsed = append(parsed, tt.Parse())
		case TokenHash:
			parsed = append(parsed, tt.Parse())
		case TokenRandomUUID:
			parsed = append(parsed, tt.Parse())
		case TokenStrFromList:
			parsed = append(parsed, tt.Parse())
		case TokenStrLower:
			parsed = append(parsed, tt.Parse())
		case TokenStrUpper:
			parsed = append(parsed, tt.Parse())
		case TokenStr:
			parsed = append(parsed, tt.Parse())
		case TokenMixLower:
			parsed = append(parsed, tt.Parse())
		case TokenMixUpper:
			parsed = append(parsed, tt.Parse())
		case TokenMix:
			parsed = append(parsed, tt.Parse())
		case Token:
			str := fmt.Sprint(tt)
			// this might be a token containing nested tokens, let's check
			for strings.Count(str, TOKEN_START) > 0 {
				if strings.Count(str, TOKEN_START) != strings.Count(str, TOKEN_END) {
					str = "" // don't include unparsed tokens!
					break
				}
				if strings.Count(str, TOKEN_START) > 1 {
					// there are nested tokens we need to solve first
					str = "[" + pp.Generate(str[1:len(str)-1]) + "]"
					continue
				}
				// this should be the outer most token, so one last parser pass
				str = pp.Generate(str)
			}
			parsed = append(parsed, str)
		case UnknownToken:
			parsed = append(parsed, "") // don't include unparsed tokens!
		default:
			parsed = append(parsed, fmt.Sprint(tt))
		}
	}
	return strings.Trim(strings.Join(parsed, ""), " \r\n")
}

func NewRandomStringGenerator(dataDir string, fileCacheErrorHandler func(err error)) *RandomStringGenerator {
	rsg := &RandomStringGenerator{
		DataDir: dataDir,
	}
	FilesCache = utils.NewFilesCache(rsg.DataDir, fileCacheErrorHandler)

	return rsg
}

// NewRandomStringGeneratorSimple returns a generator that uses $HOME/.rsg-data as its data directory
// and ignores file cache errors.
func NewRandomStringGeneratorSimple() *RandomStringGenerator {
	homedir, _ := os.UserHomeDir()
	homedir = filepath.Join(homedir, ".rsg-data")
	if err := os.MkdirAll(homedir, 0755); err != nil {
		panic("could not create data dir (" + homedir + ")")
	}
	rsg := &RandomStringGenerator{
		DataDir: homedir,
	}
	FilesCache = utils.NewFilesCache(rsg.DataDir, func(err error) { /* silently ignore errors */ })

	return rsg
}
