package main

import (
	"fmt"
	"veryhour/helper"
)

var s = `知乎知乎搜索 搜索 提问
【一】匹配 @用户名
首先分析下微博中从哪里开始到哪里@911 结束才是一个完整的用户名，按照常规的@insio 表现形式，一般是以@开头，以：结尾，中间为用户的名称。 
匹配表达式就可写为： @[^：:]+ 
这是简单的写法，但是有些是在微博之后再 ，还有就是连续@的情况，还有些是以逗号等结束的，因此完善一下 
修改为： @[^,，：:\s@]+ 
但是这些匹配都是从形式上进行了一个大概的归类，但是为了更为@严谨，就要彻底分析其用户名的具体格式，
例如@新浪微博 中的用户名格式为是“4-30个字符，支持英文、数字、"_"或减号” 
，也就是说，支持中文、字母、数字、下划线及减号，并且是4到30个字符（这里暂且认为汉字为一个字符）那么在写匹配的表达式的时候就可以这么来写: 
 
首知乎知乎搜索 搜索 提问
【一】匹配  @用户名 首先分析下微博中从哪里开始到哪里@911 结束才是一个完整的用户名，按照常规的@insio 表现形式，一般是以@开头，以：结尾，中间为用户的名称。 
匹配表达式就可写为： 这是简单的写法，但是有些是在微博之后再@的，还有就是连续@的情况，还有些是以逗号等结束的，因此完善一下 
修改为： @盈
但是这些匹配都是从形式上进行了一个大概的归类，但是为了更为@严谨，就要彻底分析其用户名的具体格式，
例如@新浪微博中的用户名格式为是“4-30个字符，支持英文、数字、"_"或减号” 
，也就是说，支持中文、字母、数字、下划线及减号，并且是4到30个字符（这里暂且认为汉字为一个字符）那么在写匹配的表达式的时候就可以这么来写: 
@[\u4e00-\u9fa5a-zA-Z0-9_-]{4,30} 

首页热门应用游戏 
还没有微博帐号？　注册|登录

文章同學http://weibo.com/wenzhang626
他还没有填写个人简介
 |北京 @李朝阳

演员文章，出演电视剧《奋斗》	
  申请认证

180
关注
41403230
粉丝
2087
微博
他的主页
微博
个人资料
关注/粉丝
相册
赞
更多 
微博
|原创 |图片
恭喜！所有的好都是你努力的结果！加油！ 
@刘欢Happy哥
辛苦媒体朋友昨天冒雨来参加答谢会，谢谢到场的所有＂欢粉＂还有为我忙前忙后的工作人员。感谢之情无语言表〜〜在此佳节之日，祝爱我的、我爱的人儿们中秋快乐！合家团员！永远快乐、幸福〜〜
       
(3100)| 转发(2573) | 评论(1847)
9月19日 14:45来自iPhone客户端
(13948)| 转发(2088)| 收藏| 评论(3671)
9月19日 14:53 来自Weico.iPhone 
不错不错！
◆
◆
@时尚抹布
在家看了一天的#小爸爸#，没吃没喝的。。。
 
(469)| 转发(1283) | 评论(242)
9月19日 06:39来自iPhone客户端
(13829)| 转发(1239)| 收藏| 评论(3848)
9月19日 06:47 来自Weico.iPhone 
意思全对 //@侯传杲: 谢谢制片人妹妹马老师弟弟文导演//@马伊琍:大飞是侯哥这辈子演得最成功的角色！因为本色出演
◆
◆
@侯传杲
钱不是问题，问题是没钱!
  
(1046)| 转发(3572) | 评论(439)
9月18日 14:06来自iPhone客户端
(9109)| 转发(1984)| 收藏| 评论(1849)
9月18日 21:29 来自Weico.iPhone 
见证历史 //@李连杰: 钱包里的公益，我们的壹基金，我们的新公益//@壹基金:我们就从这儿开始吧，日行一善，温暖每一天，温暖爱的365天，做壹基金的主人。
◆
◆
@支付宝
每1人＋每1天＋每1元＝爱的365天 【支付宝钱包下载】：d.alipay.com @支付宝公益 @壹基金 @李连杰 @陈坤 @文章同学
 
(226)| 转发(1864) | 评论(187)
9月18日 17:16来自专业版微博
(3671)| 转发(1357)| 收藏| 评论(640)
9月18日 19:47 来自Weico.iPhone 
有这奖吗？
◆
◆
@刘春
《小爸爸》里那个小娃娃太稀罕人了，就是他爹雨果太不着调了。这个片子应该入围奥斯卡电视剧奖。@文章同學 搞得不赖。
(1141)| 转发(2006) | 评论(581)
9月18日 15:29来自iPhone客户端
(9690)| 转发(1644)| 收藏| 评论(3388)
9月18日 15:51 来自Weico.iPhone |  举报
查看更多微博»
音乐
赞过(1)
 
坚持爱
艺人：文章(大陆)
风格：流行
314万人听过
(1.2万) | 分享
更多»
微关系
我们之间的共同关系
杰月映青衫

Margarita0727dalexzhan 赞同
【一】匹配 @用户名 
首先分析下微博中从哪里开始到哪里结束才是一个完整的用户名，按照常规的表现形式，一般是以@开头，以：结尾，中间为用户的名称。 
匹配表达式就可写为： @[^：:]+ 
这是简单的写法，但是有些是在微博之后再@的，还有就是连续@的情况，还有些是以逗号等结束的，因此完善一下 
修改为： @[^,，：:\s@]+ 
但是这些匹配都是从形式上进行了一个大概的归类，但是为了更为严谨，就要彻底分析其用户名的具体格式，例如新浪微博中的用户名格式为是“4-30个字符，支持英文、数字、"_"或减号” 
，也就是说，支持中文、字母、数字、下划线及减号，并且是4到30个字符（这里暂且认为汉字为一个字符）那么在写匹配的表达式的时候就可以这么来写: 
@[\u4e00-\u9fa5a-zA-Z0-9_-]{4,30} 


【二】匹配 #话题# 
这个相对相对就简单了很多，前后都是#，以#号开始并以#结束 
匹配表达式写为： #[^#]+# 

申请认证开放平台
企业微博连接网站
微博标识广告服务
微博帮助　 意见反馈　 开放平台　 微博招聘　 新浪网导航 社区管理中心　 微博社区公约
北京微梦创科网络技术有限公司 京网文[2011]0398-130号 京ICP证100780号
Copyright © 1996-2013 SINA
返回顶部
【二】匹配 #话题# 
这个相对相对就简单了很多，前后都是#，以#号开始并以#结束 
匹配表达式写为： #[^#]+# 籍？ 7 个回答
正则表达式中「正则」一词出典何处？是谁第一个在数学/计算机领域使用这个翻译的？ 5 个回答
如何避免写微博总有错别字？ 20 个回答
“抠鼻子”这个表情要表达的是什么情绪和心态？ 9 个@回答
直角引号「」比较难输入，为什么知乎用户喜欢修改问题中的双引号“”为@直角引号？ 6 个回答
「呵呵」用@日语怎么表达？ 11 个回答
「呵呵」如何用 @英语 表达？ 41 个回答
知乎 @iPhone 客户端 @koko
下载并加入知乎，随时随地提问解惑分享知识，发现更大的世界。
查看详情 »
知乎阅读 @知乎• © 2013 知乎

演员文章，出演电视剧  
|原创 |图片
恭喜！所有的好都是你努力的结果！加油！ 
@刘欢Happy哥
辛苦媒体朋友昨天冒雨来参加答谢会，谢谢到场的所有＂欢粉＂还有为我忙前忙后的工作人员。感谢之情无语言表〜〜在此佳节之日，祝爱我的、我爱的人儿们中秋快乐！合家团员！永远快乐、幸福〜〜
       
(3100)| 转发(2573 片人妹妹马老师弟弟文导演//@马伊琍:大飞是侯哥这辈子演得最成功的角色！ 
@侯传杲
钱不是问题，问题是没钱!
  
(1046)| 转发(3572) | 评论(439)
9月18日 14:06来自iPhone客户端
(9109)| 转发(1984)| 收藏| 评论(1849)
9月18日 21:29 来自Weico.iPhone 
见证历史 //@李连杰: 钱包里的公益，我们的壹基金，我们的新公益//@壹基金:我们就从这儿
开始吧，日行一善，温暖每一天，温暖爱的365天，做壹基金的主人。
 
@支付宝
每1人＋每1天＋每1元＝爱的365天 【支付宝钱包下载】：d.alipay.com @支付宝公益 @壹基金 @李连杰 @陈坤 @文章同学
 
(226)| 转发(1864) | 评论( 
有这 
@刘春
《小爸爸》里那个小娃娃太稀罕人了，就是他爹雨果太不着调了。这个片子应该入围奥斯卡电视剧奖。@文章同學 搞得不赖。
(1141)| 转发(2006) | 评论(581)
9月 
音乐
赞过(1)
  
微关系
我们之间的共同关系
他的关注(180)
更多»

大鹏董成鹏

@turbosun

林俊杰

关小悦joy
他的粉丝(41403230)
更多»

手机用户3476456250

@kanyofl

@杰月映青衫

@Mar@garita0727d
找找感兴趣的人
名人堂微博会员
微博达人 
投票微游戏
@手机 @玩微博
@WAP版 @短彩发微博
@iPhone客户端iPad客户端
@Android客户端
@认证&合作
@申请认证开放平台
@企业微博连接网站
@微博标识广告服务
@微博帮助　 @意见反馈　 @开放平台　 @微博招聘　 @新浪网导航 @社区管理中心　 @微博社区公约
北京微梦创科网络技术有限公司 京网文[2011]0398-130号 京ICP证100780号
Copyright © 1996-2013 @SINA
返回顶部
【二】匹配 #话题# 
这个相对相对就简单了很多，前后都是#，以#号开始并以#结束 
匹配表达式写为： #[^#]+# 籍？ 7 个回答
正则表达式中「正则」一词出典何处？是谁第一个在数学/计算机领域使用这个翻译的？ 5 个回答
如何避免写微博总有错别字？ 20 个回答
“抠鼻子”这个表情要表达的是什么情绪和心态？ 9 个@回答
直角引号「」比较难输入，为什么知乎用户喜欢修改问题中的双引号“”为@直角引号  6 个回答
「呵呵」用 语怎么表达？ 11 个回答
「呵呵」如何用 闻同学 英语怎么表达？ 41 个回答
知乎 @iPhone 客户端
下载并加入知乎，随时随地提问解惑分享知识，发现更大的世界。
查看详情 »
知乎阅  • © 2013 @知乎`

func main() {
	/*
		p, err := helper.GetPage("http://www.zhihu.com/question/19816153")
		fmt.Println(err)
	*/
	fmt.Println(helper.AtUsers(s))
}