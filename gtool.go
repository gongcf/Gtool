package gtool

import (
	"github.com/gongcf/gtool/compress"
	"github.com/gongcf/gtool/encrypt/aes"
	"github.com/gongcf/gtool/encrypt/base64"
	"github.com/gongcf/gtool/encrypt/hex"
	"github.com/gongcf/gtool/encrypt/md5"
	"github.com/gongcf/gtool/file"
	"github.com/gongcf/gtool/rand"
	"github.com/gongcf/gtool/response"
	"github.com/gongcf/gtool/str"
	"github.com/gongcf/gtool/system"
)

var (
	// File files
	File file.GFile
	// System utilities
	System system.GSystem
	// Rand utilities
	Rand rand.GRand
	// Ret utilities
	Ret response.GResponse
	// Str string
	Str str.Gstr
	// Compress utilities
	Compress compress.GCompress
	// Md5 utilities
	Md5 md5.GMd5
	// Aes utilities
	Aes aes.GAes
	// Base64 utilities
	Base64 base64.GBase64
	// Hex utilities
	Hex hex.GHex
)
