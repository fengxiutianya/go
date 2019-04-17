.PHONY: cleanall cleanobj cleandiff files
@## 测试
files : program *.o *.diff
	touch program;touch a.o;toucha.diff
cleanall : cleanobj cleandiff
	rm program
cleanobj :
	rm *.o
cleandiff :
	rm *.diff