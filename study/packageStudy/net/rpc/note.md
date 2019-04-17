### 只有满足了以下条件的方法才能被远程调用
>1. The method‘s type is exported
>2. 该方法是exported
>3. 方法有俩个参数哦，且必须都是exported or builtin
>4. 该方法的第二个参数必须是一个指针
>5. 方法必须返回一个error