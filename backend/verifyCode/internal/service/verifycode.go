package service

import (
	"context"
	"math/rand"
	pb "verifyCode/api/verifyCode"
)

type VerifyCodeService struct {
	pb.UnimplementedVerifyCodeServer
}

func NewVerifyCodeService() *VerifyCodeService {
	return &VerifyCodeService{}
}

func (s *VerifyCodeService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeRequest) (*pb.GetVerifyCodeReply, error) {

	return &pb.GetVerifyCodeReply{
		Code: RandCode(int(req.Length), req.Type),
	}, nil
}

func RandCode(l int, tp pb.TYPE) string {

	switch tp {
	case pb.TYPE_DIGIT:
		return randCode("0123456789", l, 4)
	case pb.TYPE_LETTER:
		return randCode("abcdefghijklmnopqrstuvwxyz", l, 5)
	case pb.TYPE_MIXED:
		return randCode("0123456789abcdefghijklmnopqrstuvwxyz", l, 6)
	default:
		return ""
	}
}

// 简单实现
//func randCode(chars string, l int) string {
//	length := len(chars)
//	result := make([]byte, l)
//
//	for i := 0; i < l; i++ {
//		index := rand.Intn(length)
//		result[i] = chars[index]
//	}
//
//	return string(result)
//}

// 优化 随机
func randCode(chars string, l, idx int) string {

	result := make([]byte, l)

	// 形成掩码
	mask := 1<<idx - 1 // 00001 000000 - 1 = 00000 111111
	idMax := 63 / idx

	// remian 还可以用几次
	// cache 生成的随机 数
	// i 长度
	for i, cache, remain := 0, rand.Int63(), idMax; i < l; {
		if remain == 0 {
			cache, remain = rand.Int63(), idMax
		}

		// 利用掩码获取随机数位
		if randIndex := int(cache & int64(mask)); randIndex < len(chars) {
			result[i] = chars[randIndex]
			i++
		}

		cache >>= idx
		idMax--

	}

	return string(result)
}
