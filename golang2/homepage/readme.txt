			《 homepage 参考文档 》
BY： 冼锦荣

===================== Angular ===========================
			
重要变量和数据：
more_button_show ： string		//文章尾部按钮显示的文章，可以是'加载更多;和'没有了'.
artical_num : 	 number			//当前首页文章的数量，限定加载文章的数量，优先加载最新文章，且数量限定在100以内，超过后more_button_show变成“没有了”



函数：
			< homepage.component.ts >
									
# get_artical()					//加载首页文章列表，数据包括图片链接，文件链接，文章简介，每次加载十个，调用 golang.service.ts 的 GetArtical(index:number) 方法

# get_head_new()				//加载首页新闻数据，共五条。调用 golang.service.ts 的 GetHeadNews() 方法

 
			< golang.service.ts >


# GetArtical(index:number)		//向服务器请求首页文章数据

# GetHeadNews()					//向服务器请求首页头部滑动图片链接数据


备注：
1.homepage.component.ts 使用到 Jquery，都放在 ngoninit() 函数中，都是实现图片轮播功能的代码		
2.首页头部的轮播图片框，图片数量暂时固定为5个，请求时将获得最新的五个数据	
	
	
=================== Golang ===============================
路由函数：

# GetImages(...)				//返回图片，例子：url=”http://129.204.193.192:4400/get?tag=&name=$.jpg“，那么将获取 路径 + tag+ name 即”1.jpg",并返回图片字节


# GetHomePageArtical(...)		//返回首页文章的数据，切片类型，post请求，获得一个整数N代表已有前面几条数据，回应 n~n+10 的文章数据  


# GetHomePageHotnews(...)		//返回首页头部滚动播放图片框需要的数据，固定返回最新的五条数据

数据库操作函数：

# GetHomePageHotnewDate(...)	//从数据库获取首页图片轮播框的所需数据，被 GetHomePageHotnews(...)调用

# GetHPADate(...)				//GetHeadpageArticalDate 的意思，被GetHomePageArtical(...)调用

# Images(...)					//从本地文件获取得图片，由GetImages(...)调用

主要的数据结构：

type HomePageNews struct{					//返回轮播图片框需要的数据
	Img_url	string 			`json:"imgurl"`
	Link_url string 		`json:"linkurl"`
	Title	string 			`json:"title"`
}

type ArticlaBox struct{						//返回文章框需要的数据
	Img_url	string 			`json:"imgurl"`
	Link_url string 		`json:"linkurl"`
	Brif	string 			`json:"brief"`
	Date	string 			`json:"date"`
}


============================== SQL =============================
创建table的命令：

1：滚动播放框的数据
create table homepagenews(
imgurl varchar(200),			//图片链接
linkurl varchar(200),			//文章链接
title varchar(200),				//新闻标题
id serial);

2：首页文章的数据
create table homepageartical (
id serial,						
imgurl varchar(200),			
linkurl varchar(200),			//文章的连接，
brif    varchar(200),			//显示在文章框的主要文字，可以是文章前面的一段话
date date);						//上传时间



=================================================================



