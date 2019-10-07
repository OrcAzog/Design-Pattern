package main

import "sync"

var once sync.Once

type  singleton struct{

}
var instance *singleton

func getInstance() *singleton{
	once.Do(func(){
		instance=new(singleton)

	})
	return instance
}
