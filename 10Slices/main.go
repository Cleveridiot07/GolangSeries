package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices Introduction");

	var fruitList =  []string{"apple", "orange"}
	// fmt.Printf("Type fruitlist %T\n", fruitList);
	// fmt.Println(fruitList);

	// we can add as many value as we want

	fruitList = append(fruitList, "mango","orange");
	// fmt.Println(fruitList);


	fruitList = append(fruitList[1:]);
	// fmt.Println(fruitList);

	fruitList = append(fruitList[:2])
	// fmt.Println(fruitList);

// datatype and then size
	highScore := make([]int,4)
	highScore[0] = 875
	highScore[1] = 900
	highScore[2] = 745
	highScore[3] = 550

	// fmt.Println(highScore);

	highScore = append(highScore, 584,748,964)
	// will reallocate the memory
	// fmt.Println(highScore);


	sort.Ints(highScore);
	// fmt.Println(highScore);

	// fmt.Println((sort.IntsAreSorted(highScore)));


	// remove a value from slice based on index
	var courses = []string{"java","python","sql","swift"};
	fmt.Println(courses);

	var index int = 2;
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses);
}