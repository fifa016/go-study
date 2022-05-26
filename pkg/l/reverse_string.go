/*
 * @Descripttion: 
 * @Author: jzh
 * @Date: 2022-05-26 16:14:11
 */
package l



func reverseString(s []byte)  {

	for i := 0; i < len(s) / 2; i++ {
		s[i], s[len(s) - 1 - i] = s[len(s) - 1 - i] , s[i]
	}
    
}