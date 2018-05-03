package main

import "fmt"

type Student struct {
	name         string
	age          int
	sex          string
	englishScore float64
	mathScore    float64
	chineseScore float64
}

func (s Student) getSum() float64 {
	sum:= s.chineseScore+s.mathScore+s.englishScore
	return sum
}

func (s Student) getAvg()  float64{
	sum := s.getSum()// 调用上面已经实现了的本类的方法，求和
	avg := sum / 3
	return avg
}

func (s Student) printInfo()  { //s ---> s1
	fmt.Printf("姓名：%s,年龄：%d，性别：%s\n",s.name,s.age,s.sex)
	fmt.Printf("\t英语成绩：%.2f,数学成绩：%.2f,语文成绩：%.2f\n",s.englishScore,s.mathScore,s.chineseScore)
	fmt.Printf("\t总分：%.2f,平均分：%.2f\n",s.getSum(),s.getAvg())
}

func main()  {
	s1 := Student{"王二狗",18,"男",68.6,99.6,58.4}
	s1.printInfo()

	//s1.getSum()

}


