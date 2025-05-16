package codetop

func addStrings(num1 string, num2 string) string {
	var res []byte
	l1 := len(num1)
	l2 := len(num2)
	// 确保num1的长度等于num2
	// 不够进行0的填充
	if l1 < l2 {
		for i := 0; i < l2-l1; i++ {
			num1 = "0" + num1
		}
	} else {
		for i := 0; i < l1-l2; i++ {
			num2 = "0" + num2
		}
	}
	l := len(num1)
	res = make([]byte, l)

	carry := 0
	for i := l - 1; i >= 0; i-- {
		// 获取当前位的值
		cur1 := int(num1[i] - '0')
		cur2 := int(num2[i] - '0')
		if cur1+cur2+carry >= 10 {
			res[i] = byte(cur1+cur2+carry-10) + '0'
			carry = 1
		} else {
			res[i] = byte(cur1+cur2+carry) + '0'
			carry = 0
		}
	}
	if carry == 1 {
		res = append([]byte{'1'}, res...)
	}
	return string(res)
}
