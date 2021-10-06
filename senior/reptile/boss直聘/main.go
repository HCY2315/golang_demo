package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	parentChildSelector() // 指定标签下的内容
	/*
		parentChildSelector()         // 指定标签下的内容
		classSelector(webcontent)     // 通过 class 获取标签数
		elementIdSelector(webcontent) // 通过 div 后面的 ID 获取数据
		elementSelector(webcontent)   // 获取所有 div 标签的内容
	*/
}

// classSelector 通过 class 获取标签数据
func classSelector(webcontent string) {
	// 筛选出 class 为 name 的 所有 元素
	selectorContent1 := ".job-name"
	goquerySelector(webcontent, selectorContent1)

	// 筛选出 class 为 abc 的 div 元素
	selectorContent2 := "div[class=job-title]"
	goquerySelector(webcontent, selectorContent2)

	// 筛选出 class 为 job-name 的 span 元素
	selectorContent3 := "span[class=job-name]"
	goquerySelector(webcontent, selectorContent3)

	// 筛选出 有 class 属性的 span 元素
	selectorContent4 := "span[class]"
	goquerySelector(webcontent, selectorContent4)
}

// 指定标签下的内容
func parentChildSelector() {
	webcontent :=
		`<body>
			<div class="aaa">DIV1</div>
			<div>DIV2</div>
			<div>DIV3</div>
			<span>
				<div>DIV4</div>
			</span>
		</body>
`
	// 获取子集中class为 aaa 的标签内容
	selectorContent1 := "body>div[class=aaa]"
	goquerySelector(webcontent, selectorContent1)

	// 获取body下所有子集
	selectorContent2 := "body>div"
	goquerySelector(webcontent, selectorContent2)

	// 获取body下所有div集
	selectorContent3 := "body div"
	goquerySelector(webcontent, selectorContent3)
}

// elementIdSelector 通过 ID 获取所有标签的数据
func elementIdSelector(webcontent string) {
	selectorContent1 := "div#filter-box"
	goquerySelector(webcontent, selectorContent1)

	// canvas id="sign-side" 演示 map 使用，一般不打开
	// selectorContent2 := "canvas#sign-side"
	// result := goquerySelector(webcontent, selectorContent2)
	// fmt.Println(result)
}

// elementSelector 获取所有 div 标签的内容
func elementSelector(webcontent string) {
	selectorContent := "div"
	goquerySelector(webcontent, selectorContent)
}

// goquerySelector 封装 爬虫函数
func goquerySelector(htmlContent string, selectorContent string) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal("err:, ", err)
		return
	}

	// selectorContent 为查询的语法
	dom.Find(selectorContent).Each(func(i int, selector *goquery.Selection) {
		fmt.Println(selector.Text())
	})

	// return dom.Find(selectorContent).Map(func(i int, s *goquery.Selection) string {
	// 	fmt.Println(s.Text())
	// 	return s.Text()
	// })
}

var webcontent string = `<!DOCTYPE html>
<!--[if lt IE 7 ]><html class="ie ie6"><![endif]-->
<!--[if IE 7 ]><html class="ie ie7"><![endif]-->
<!--[if IE 8 ]><html class="ie ie8"><![endif]-->
<!--[if IE 9 ]><html class="ie9"><![endif]-->
<!--[if (gt IE 9)|!(IE)]><!--><html class="standard"><!--<![endif]-->
<head>
    <meta charset="utf-8">
    <title>「北京golang后端开发招聘」-2021年北京golang后端开发人才招聘信息 - BOSS直聘</title>
    <meta name="keywords" content="北京golang后端开发招聘,北京golang后端开发人才网,北京golang后端开发招聘网"/>
    <meta name="description" content="BOSS直聘为求职者提供2021年北京golang后端开发招聘信息，百万Boss在线直聘，直接开聊，在线面试，找工作就上BOSS直聘网站或APP，直接与Boss开聊吧！"/>
    <link href="https://static.zhipin.com/zhipin-geek/v463/web/geek/css/main.css" type="text/css" rel="stylesheet">
    <link rel="canonical" href="https://www.zhipin.com/c101010100-p100199/?query=golang">
    <meta name="applicable-device" content="pc">
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1"/>
    <link rel="alternate" media="only screen and(max-width: 640px)" href="https://m.zhipin.com/c101010100-p100199/?query=golang">
    <script>
        var positionList = [];
        var _STGY = {
            login_source:2
        };
    </script>
</head>
<body>
<div id="wrap" class="search-job-list-wrap">
<script>
    var staticPath = 'https://static.zhipin.com/v2';
    _PAGE = {
        checkMobileUrl: "/registe/sendSms.json",
        regMobileUrl: "/registe/save.json",
        loginMobileUrl: "/login/phone.json",
        loginAccountUrl: "/login/account.json",
        getRandomKeyUrl: "/captcha/randkey.json",
        verifyImgUrl: "/captcha/?randomKey={randomKey}",
        getPositionUrl: "/user/position.json",
        citySiteName:"北京站",
        citySiteCode:"101010100",
        skillsUrl: "/common/data/positionSkill"
    }
</script>
<div id="header">
    <div class="inner home-inner">
        <div class="logo">
            <a href="https://www.zhipin.com/" ka="header-home-logo" title="BOSS直聘"><span>BOSS直聘</span></a>
        </div>
        <div class="nav">
            <ul>
                <li class=""><a ka="header-home" href="https://www.zhipin.com/beijing/">首页</a></li>
                <li class="cur"><a ka="header-job" href="https://www.zhipin.com/job_detail/">职位</a></li>
                <li class=""><a class="nav-school-new" ka="tab_school_recruit_click" href="https://www.zhipin.com/xiaoyuan/">校园</a></li>
                <li class=""><a class="nav-overseas-new" ka="tab_overseas_click" data-overseas-ip="false"  href="https://www.zhipin.com/returnee_jobs/">海归</a></li>
                <li class=""><a ka="header_brand" href="https://www.zhipin.com/gongsi/">公司</a></li>
                <li class=""><a ka="header-app" href="https://app.zhipin.com/">APP</a></li>
                <li class=""><a ka="header-article" href="https://news.zhipin.com/">资讯</a></li>
                <li class=""><a class="nav-find" ka="header-find" href="https://www.zhipin.com/faxian/">发现</a></li>
            </ul>
        </div>
        <div class="user-nav">
                <!--未登录-->
                <div class="btns">
                    <a href="https://www.zhipin.com/wapi/zpgeek/resume/parser.html" rel="nofollow" ka="nlp_resume_upload" class="link-sign-resume" title="上传简历，解析内容，完善注册">上传简历</a>
                    <a href="https://www.zhipin.com/wapi/zpgeek/resume/parser.html" rel="nofollow" ka="header-geek" class="link-sign-resume" title="上传简历，解析内容，完善注册">我要找工作<span class="new" style="display: inline;">hot</span></a>
                    <a href="https://signup.zhipin.com?intent=1" rel="nofollow" class="link-scan" ka="header-boss" title="我要招聘">我要招聘</a>
                    <a href="https://signup.zhipin.com" ka="header-register" class="btn btn-outline">注册</a>
                    <a href="https://login.zhipin.com" ka="header-login" class="btn btn-outline">登录</a>
                </div>
        </div>
    </div>
</div>    <div id="filter-box">
        <div class="inner home-inner">
            <div class="search-box search-box-new clear-fix">
                <div class="search-form ">
                    <form action="/job_detail/" method="get">
                        <div class="search-form-con">
                            <div class="city-sel">
                                <i class="line"></i>
                                <span class="label-text"><b>北京</b><i class="icon-arrow-down"></i></span>
                            </div>
                            <div class="industry-sel" ka="search_bos_sel_industry">
                                <i class="line"></i>
                                <span class="label-text"><b>公司行业</b><i class="icon-arrow-down"></i></span>
                            </div>
                            <div class="position-sel" ka="search_box_sel_jobtype">
                                <i class="line"></i>
                                <span class="label-text"><b>后端开发</b><i class="icon-arrow-down"></i></span>
                            </div>
                            <p class="ipt-wrap"><input type="text" name="query" ka="search-keyword" value="golang" autocomplete="off" class="ipt-search" maxlength="50" placeholder="搜索职位、公司"></p>
                        </div>
                        <input type="hidden" name="city" class="city-code" value="101010100" />
                        <input type="hidden" name="industry" class="industry-code" value="" />
                        <input type="hidden" name="position" class="position-code" value="100199" />
                        <button class="btn btn-search">搜索</button>
                        <div class="suggest-result">
                            <ul>
                            </ul>
                        </div>
                        <div class="city-box">
                            <ul class="dorpdown-province">
                            </ul>
                            <div class="dorpdown-city">
                            </div>
                        </div>
                        <div class="industry-box">
                            <ul>
                            </ul>
                        </div>
                        <div class="position-box">
                            <div class="dropdown-menu">
                                <script>
                                    var jobData = positionList;
                                </script>
                                <div class="select-tree" data-level="3"></div>
                            </div>
                        </div>
                    </form>
                </div>
                <!-- 广告 -->
                <!--
                <div class="search-banner">
                    <img src="http://172.16.24.210/v2/web/geek/images/listbanner.png" alt="">
                </div>
                -->
                    <div class="search-box-login">
                        <a href="javascript:void(0);" ka="cpc_job_list_signup"><span class="search-box-login-icon"></span>登录，查看更多职位<span class="search-box-login-close" ka="cpc_job_list_signup_cancel"></span></a>
                    </div>
            </div>
            <div class="condition-box">
                <dl class="condition-city show-condition-district">
                    <dd class="city-wrapper">
                        <a href="javascript:;" class="disabled" title="北京后端开发招聘">北京</a>
                            <em class="icon-arrow-right"></em>
                            <a href="javascript:;" class='link-district selected'>
                                不限</a>
                            <em class="icon-arrow-right"></em>
                        <span class="hot-text">热门城市：</span>
                                <a href="/c100010000-p100199/?query=golang" ka="sel-city-100010000">全国</a>
                                <a href="/c101010100-p100199/?query=golang" ka="sel-city-101010100">北京</a>
                                <a href="/c101020100-p100199/?query=golang" ka="sel-city-101020100">上海</a>
                                <a href="/c101280100-p100199/?query=golang" ka="sel-city-101280100">广州</a>
                                <a href="/c101280600-p100199/?query=golang" ka="sel-city-101280600">深圳</a>
                                <a href="/c101210100-p100199/?query=golang" ka="sel-city-101210100">杭州</a>
                                <a href="/c101030100-p100199/?query=golang" ka="sel-city-101030100">天津</a>
                                <a href="/c101110100-p100199/?query=golang" ka="sel-city-101110100">西安</a>
                                <a href="/c101190400-p100199/?query=golang" ka="sel-city-101190400">苏州</a>
                                <a href="/c101200100-p100199/?query=golang" ka="sel-city-101200100">武汉</a>
                                <a href="/c101230200-p100199/?query=golang" ka="sel-city-101230200">厦门</a>
                                <a href="/c101250100-p100199/?query=golang" ka="sel-city-101250100">长沙</a>
                                <a href="/c101270100-p100199/?query=golang" ka="sel-city-101270100">成都</a>
                                <a href="/c101180100-p100199/?query=golang" ka="sel-city-101180100">郑州</a>
                                <a href="/c101040100-p100199/?query=golang" ka="sel-city-101040100">重庆</a>
                        <a href="javascript:;" class="btn-allcity">全部城市</a>
                    </dd>
                </dl>
                    <dl class='condition-district  show-condition-district '>
                        <dd>
                            <a  class="selected"  href="/c101010100-p100199/?query=golang" ka="sel-business-0">不限</a>
                                    <a                                              href="/c101010100-p100199/b_%E6%B5%B7%E6%B7%80%E5%8C%BA/?query=golang"
                                        
                                        ka="sel-business-0">海淀区</a>
                                    <a                                              href="/c101010100-p100199/b_%E6%9C%9D%E9%98%B3%E5%8C%BA/?query=golang"
                                        
                                        ka="sel-business-1">朝阳区</a>
                                    <a                                              href="/c101010100-p100199/b_%E4%B8%9C%E5%9F%8E%E5%8C%BA/?query=golang"
                                        
                                        ka="sel-business-2">东城区</a>
                                    <a                                              href="/c101010100-p100199/b_%E6%98%8C%E5%B9%B3%E5%8C%BA/?query=golang"
                                        
                                        ka="sel-business-3">昌平区</a>
                                    <a                                              href="/c101010100-p100199/b_%E8%A5%BF%E5%9F%8E%E5%8C%BA/?query=golang"
                                        
                                        ka="sel-business-4">西城区</a>
                                    <a                                              href="/c101010100-p100199/b_%E5%A4%A7%E5%85%B4%E5%8C%BA/?query=golang"
                                        
                                        ka="sel-business-5">大兴区</a>
                                    <a                                              href="/c101010100-p100199/b_%E4%B8%B0%E5%8F%B0%E5%8C%BA/?query=golang"
                                        
                                        ka="sel-business-6">丰台区</a>
                                    <a                                              href="/c101010100-p100199/b_%E7%9F%B3%E6%99%AF%E5%B1%B1%E5%8C%BA/?query=golang"
                                        
                                        ka="sel-business-7">石景山区</a>
                                    <a                                              href="/c101010100-p100199/b_%E9%A1%BA%E4%B9%89%E5%8C%BA/?query=golang"
                                        
                                        ka="sel-business-8">顺义区</a>
                                    <a                                              href="/c101010100-p100199/b_%E9%80%9A%E5%B7%9E%E5%8C%BA/?query=golang"
                                        
                                        ka="sel-business-9">通州区</a>
                                    <a                                              href="/c101010100-p100199/b_%E6%88%BF%E5%B1%B1%E5%8C%BA/?query=golang"
                                        
                                        ka="sel-business-10">房山区</a>
                                    <a  
                                            href="javascript:;" class="btn-disabled"
                                        ka="sel-business-11">平谷区</a>
                                    <a  
                                            href="javascript:;" class="btn-disabled"
                                        ka="sel-business-12">怀柔区</a>
                                    <a  
                                            href="javascript:;" class="btn-disabled"
                                        ka="sel-business-13">延庆区</a>
                                    <a  
                                            href="javascript:;" class="btn-disabled"
                                        ka="sel-business-14">密云区</a>
                                    <a  
                                            href="javascript:;" class="btn-disabled"
                                        ka="sel-business-15">门头沟区</a>
                        </dd>
                    </dl>
                    <dl class='condition-area '>
                        <dd>
                            <a  class="selected"  href="/c101010100-p100199/?query=golang" ka="sel-area-0">不限</a>
                        </dd>
                    </dl>
            </div>
            <div class="box-shadow-box"></div>
            <div class="filter-select-box">
                <div class="dropdown-wrap">
                    
                    <span class="dropdown-select">
                            <input type="text" ka="sel-exp" value="工作经验" class="ipt" readonly>
                        <i class="icon-select-arrow"></i>
                        <div class="dropdown-menu">
                            <ul>
                                            <li><a href="/c101010100-p100199/?query=golang" ka="sel-exp-0" rel="nofollow">不限</a></li>
                                            <li><a href="/c101010100-p100199/e_108/?query=golang" ka="sel-exp-108" rel="nofollow">在校生</a></li>
                                            <li><a href="/c101010100-p100199/e_102/?query=golang" ka="sel-exp-102" rel="nofollow">应届生</a></li>
                                            <li><a href="/c101010100-p100199/e_103/?query=golang" ka="sel-exp-103" rel="nofollow">1年以内</a></li>
                                            <li><a href="/c101010100-p100199/e_104/?query=golang" ka="sel-exp-104" rel="nofollow">1-3年</a></li>
                                            <li><a href="/c101010100-p100199/e_105/?query=golang" ka="sel-exp-105" rel="nofollow">3-5年</a></li>
                                            <li><a href="/c101010100-p100199/e_106/?query=golang" ka="sel-exp-106" rel="nofollow">5-10年</a></li>
                                            <li><a href="/c101010100-p100199/e_107/?query=golang" ka="sel-exp-107" rel="nofollow">10年以上</a></li>
                            </ul>
                        </div>
                    </span>
                </div>
                <div class="dropdown-wrap">
                    
                    <span class="dropdown-select">
                            <input type="text" value="学历要求" ka="sel-degree" class="ipt" readonly>
                        <i class="icon-select-arrow"></i>
                        <div class="dropdown-menu">
                            <ul>
                                        <li><a href="/c101010100-p100199/?query=golang" ka="sel-degree-0" rel="nofollow">不限</a></li>
                                        <li><a href="/c101010100-p100199/d_209/?query=golang" ka="sel-degree-209" rel="nofollow">初中及以下</a></li>
                                        <li><a href="/c101010100-p100199/d_208/?query=golang" ka="sel-degree-208" rel="nofollow">中专/中技</a></li>
                                        <li><a href="/c101010100-p100199/d_206/?query=golang" ka="sel-degree-206" rel="nofollow">高中</a></li>
                                        <li><a href="/c101010100-p100199/d_202/?query=golang" ka="sel-degree-202" rel="nofollow">大专</a></li>
                                        <li><a href="/c101010100-p100199/d_203/?query=golang" ka="sel-degree-203" rel="nofollow">本科</a></li>
                                        <li><a href="/c101010100-p100199/d_204/?query=golang" ka="sel-degree-204" rel="nofollow">硕士</a></li>
                                        <li><a href="/c101010100-p100199/d_205/?query=golang" ka="sel-degree-205" rel="nofollow">博士</a></li>
                            </ul>
                        </div>
                    </span>
                </div>
                <div class="dropdown-wrap">
                    
                    <span class="dropdown-select">
                            <input type="text" ka="sel-salary" value="薪资要求" class="ipt" readonly>
                        <i class="icon-select-arrow"></i>
                        <div class="dropdown-menu">
                            <ul>
                                        <li><a href="/c101010100-p100199/?query=golang" ka="sel-salary-0" rel="nofollow">不限</a></li>
                                        <li><a href="/c101010100-p100199/y_1/?query=golang" ka="sel-salary-1" rel="nofollow">3K以下</a></li>
                                        <li><a href="/c101010100-p100199/y_2/?query=golang" ka="sel-salary-2" rel="nofollow">3-5K</a></li>
                                        <li><a href="/c101010100-p100199/y_3/?query=golang" ka="sel-salary-3" rel="nofollow">5-10K</a></li>
                                        <li><a href="/c101010100-p100199/y_4/?query=golang" ka="sel-salary-4" rel="nofollow">10-15K</a></li>
                                        <li><a href="/c101010100-p100199/y_5/?query=golang" ka="sel-salary-5" rel="nofollow">15-20K</a></li>
                                        <li><a href="/c101010100-p100199/y_6/?query=golang" ka="sel-salary-6" rel="nofollow">20-30K</a></li>
                                        <li><a href="/c101010100-p100199/y_7/?query=golang" ka="sel-salary-7" rel="nofollow">30-50K</a></li>
                                        <li><a href="/c101010100-p100199/y_8/?query=golang" ka="sel-salary-8" rel="nofollow">50K以上</a></li>
                            </ul>
                        </div>
                    </span>
                </div>
                <div class="dropdown-wrap">
                    
                    <span class="dropdown-select">
                            <input type="text" ka="sel-stage" value="融资阶段" class="ipt" readonly>
                        <i class="icon-select-arrow"></i>
                        <div class="dropdown-menu">
                            <ul>
                                        <li><a href="/c101010100-p100199/?query=golang" ka="sel-stage-0" rel="nofollow">不限</a></li>
                                        <li><a href="/c101010100-p100199/t_801/?query=golang" ka="sel-stage-801" rel="nofollow">未融资</a></li>
                                        <li><a href="/c101010100-p100199/t_802/?query=golang" ka="sel-stage-802" rel="nofollow">天使轮</a></li>
                                        <li><a href="/c101010100-p100199/t_803/?query=golang" ka="sel-stage-803" rel="nofollow">A轮</a></li>
                                        <li><a href="/c101010100-p100199/t_804/?query=golang" ka="sel-stage-804" rel="nofollow">B轮</a></li>
                                        <li><a href="/c101010100-p100199/t_805/?query=golang" ka="sel-stage-805" rel="nofollow">C轮</a></li>
                                        <li><a href="/c101010100-p100199/t_806/?query=golang" ka="sel-stage-806" rel="nofollow">D轮及以上</a></li>
                                        <li><a href="/c101010100-p100199/t_807/?query=golang" ka="sel-stage-807" rel="nofollow">已上市</a></li>
                                        <li><a href="/c101010100-p100199/t_808/?query=golang" ka="sel-stage-808" rel="nofollow">不需要融资</a></li>
                            </ul>
                        </div>
                    </span>
                </div>
                <div class="dropdown-wrap">
                    <span class="dropdown-select">
                            <input type="text" ka="sel-scale" value="公司规模" class="ipt" readonly>
                        <i class="icon-select-arrow"></i>
                        <div class="dropdown-menu">
                            <ul>
                                       <li><a href="/c101010100-p100199/?query=golang" ka="sel-scale-0" rel="nofollow">不限</a></li>
                                       <li><a href="/c101010100-p100199/s_301/?query=golang" ka="sel-scale-301" rel="nofollow">0-20人</a></li>
                                       <li><a href="/c101010100-p100199/s_302/?query=golang" ka="sel-scale-302" rel="nofollow">20-99人</a></li>
                                       <li><a href="/c101010100-p100199/s_303/?query=golang" ka="sel-scale-303" rel="nofollow">100-499人</a></li>
                                       <li><a href="/c101010100-p100199/s_304/?query=golang" ka="sel-scale-304" rel="nofollow">500-999人</a></li>
                                       <li><a href="/c101010100-p100199/s_305/?query=golang" ka="sel-scale-305" rel="nofollow">1000-9999人</a></li>
                                       <li><a href="/c101010100-p100199/s_306/?query=golang" ka="sel-scale-306" rel="nofollow">10000人以上</a></li>
                            </ul>
                        </div>
                    </span>
                </div>
                <a href="/c101010100/?query=golang" ka="empty-filter" class="empty-filter" rel="nofollow">清空筛选条件</a>
            </div>
        </div>
    </div>
    <div id="main" class="inner home-inner">
        <div class="job-box">
            <div class="sider">

    <div class="sign-wrap sign-wrap-v2">
        <div class="sign-form sign-quick" style="display: block;">
            <canvas id="sign-side" width="250" height="50"></canvas>
            <h3 class="title">各大行业职位任你选</h3>
            <form action="/wapi/zppassport/login/phone" method="post" ka="sup_g_right_done">
                <div class="tip-error tip-error-form"></div>
                <div class="form-row row-select">
                    <span class="dropdown-select"><i class="icon-select-arrow"></i><em class="text-select">+86</em><input type="hidden" name="regionCode" value="+86"/></span>
                    <span class="ipt-wrap"><i class="icon-sign-phone"></i><input type="tel" class="ipt ipt-phone required" ka="signin-account" placeholder="手机号" name="phone"/></span>
                    <div class="dropdown-menu">
<ul>
    <li data-val="+86"><span class="num">+86</span>中国大陆</li>
    <li data-val="+852"><span class="num">+852</span>中国香港</li>
    <li data-val="+853"><span class="num">+853</span>中国澳门</li>
    <li data-val="+886"><span class="num">+886</span>中国台湾</li>
    <li data-val="+1"><span class="num">+1</span>美国</li>
    <li data-val="+81"><span class="num">+81</span>日本</li>
    <li data-val="+44"><span class="num">+44</span>英国</li>
    <li data-val="+82"><span class="num">+82</span>韩国</li>
    <li data-val="+33"><span class="num">+33</span>法国</li>
    <li data-val="+7"><span class="num">+7</span>俄罗斯</li>
    <li data-val="+39"><span class="num">+39</span>意大利</li>
    <li data-val="+65"><span class="num">+65</span>新加坡</li>
    <li data-val="+63"><span class="num">+63</span>菲律宾</li>
    <li data-val="+60"><span class="num">+60</span>马来西亚</li>
    <li data-val="+64"><span class="num">+64</span>新西兰</li>
    <li data-val="+34"><span class="num">+34</span>西班牙</li>
    <li data-val="+95"><span class="num">+95</span>缅甸</li>
    <li data-val="+230"><span class="num">+230</span>毛里求斯</li>
    <li data-val="+263"><span class="num">+263</span>津巴布韦</li>
    <li data-val="+20"><span class="num">+20</span>埃及</li>
    <li data-val="+92"><span class="num">+92</span>巴基斯坦</li>
    <li data-val="+880"><span class="num">+880</span>孟加拉国</li>
</ul>                    </div>
                    <div class="tip-error"></div>
                </div>
                <div class="form-row">
                    <div class="row-code" id="loginVerrifyCode2"></div>
                    <div class="tip-error"></div>
                </div>
                <input type="hidden" name="csessionId">
                <input type="hidden" name="csig">
                <input type="hidden" name="ctoken">
                <input type="hidden" name="cscene" value="nc_login">
                <input type="hidden" value="FFFF0N00000000006DC1" name="cappKey">
                <div class="form-row">
                    <span class="ipt-wrap">
                        <i class="icon-sign-sms"></i>
                        <input type="text" class="ipt ipt-sms required" ka="signup-sms" placeholder="短信验证码" name="phoneCode" maxlength="6"/>
                        <input type="hidden" name="smsType" value="7"/>
                        <button type="button" class="btn btn-sms" data-url="/wapi/zppassport/send/smsCode">发送验证码</button>
                    </span>
                    <div class="tip-error"></div>
                </div>
                <div class="form-btn">
                    <button type="submit" class="btn" ka="sup_g_right">登录/注册</button>
                </div>
                <div class="text-tip">
                  <div class="tip-error"></div>
                  <input type="checkbox" class="agree-policy" name="policy">同意BOSS直聘<a href="https://about.zhipin.com/agreement?id=registerprotocol" ka="link-privacy1" class="user-agreement" target="_blank">《用户协议》</a>
                  <a href="https://about.zhipin.com/agreement?id=personalinfopro" ka="link-privacy2" class="user-agreement" target="_blank">《个人信息保护政策》</a><br>
                </div>
            </form>
        </div>
    </div>

                    <div class="complete-resume">
                        <a href="https://www.zhipin.com/wapi/zpgeek/resume/parser.html" class="btn" rel="nofollow">上传简历一键注册</a>
                    </div>
                <div class="sider-list" style="display:none">
                    <h3>看过的职位</h3>
                    <ul>
                    </ul>
                </div>
                   <div class="sider-list about-search">
                       <h3>相关搜索</h3>
                       <ul>
                              <li><a href="/job_detail/?query=go%E8%AF%AD%E8%A8%80%E5%BC%80%E5%8F%91%E5%B7%A5%E7%A8%8B%E5%B8%88" target="_blank" ka="related-search-0">go语言开发工程师<span class="icon-arrow-right"></span></a></li>
                              <li><a href="/job_detail/?query=golang%E5%BC%80%E5%8F%91%E5%B7%A5%E7%A8%8B%E5%B8%88" target="_blank" ka="related-search-1">golang开发工程师<span class="icon-arrow-right"></span></a></li>
                              <li><a href="/job_detail/?query=erlang" target="_blank" ka="related-search-2">erlang<span class="icon-arrow-right"></span></a></li>
                              <li><a href="/job_detail/?query=go%E5%BE%AE%E6%9C%8D%E5%8A%A1" target="_blank" ka="related-search-3">go微服务<span class="icon-arrow-right"></span></a></li>
                              <li><a href="/job_detail/?query=golang%E5%AE%9E%E4%B9%A0" target="_blank" ka="related-search-4">golang实习<span class="icon-arrow-right"></span></a></li>
                              <li><a href="/job_detail/?query=go%E6%B8%B8%E6%88%8F" target="_blank" ka="related-search-5">go游戏<span class="icon-arrow-right"></span></a></li>
                              <li><a href="/job_detail/?query=go%E5%AE%9E%E4%B9%A0%E7%94%9F" target="_blank" ka="related-search-6">go实习生<span class="icon-arrow-right"></span></a></li>
                              <li><a href="/job_detail/?query=haskell" target="_blank" ka="related-search-7">haskell<span class="icon-arrow-right"></span></a></li>
                              <li><a href="/job_detail/?query=go%E5%BC%80%E5%8F%91%E5%B7%A5%E7%A8%8B%E5%B8%88" target="_blank" ka="related-search-8">go开发工程师<span class="icon-arrow-right"></span></a></li>
                              <li><a href="/job_detail/?query=c%2B%2B%E6%9C%8D%E5%8A%A1%E5%99%A8" target="_blank" ka="related-search-9">c++服务器<span class="icon-arrow-right"></span></a></li>
                       </ul>
                   </div>
                <!--<div class="promotion-img nomargin"><a href="/activity/personality/index.html" target="_blank" ka="ad_banner_0"><img src="https://z.zhipin.com/web/upload/market/coop/20180421-03.jpg" alt="" /></a></div> -->
                <div class="promotion-img"><a href="/app.html" target="_blank" ka="ad_banner_1" rel="nofollow"><img src="https://static.zhipin.com/zhipin-geek/v463/web/geek/images/pro-1.png" alt="" /></a></div>
                <div class="promotion-img"><a href="/user/login.html?initType=3" target="_blank" ka="ad_banner_2" rel="nofollow"><img src="https://static.zhipin.com/zhipin-geek/v463/web/geek/images/pro-2.png" alt="" /></a></div>
                <div class="promotion-img"><a href="/salaryxc/" target="_blank" ka="query_pay_click"><img src="https://static.zhipin.com/zhipin-geek/v463/web/geek/images/pro-3.jpg" alt="" /></a></div>
            </div>
                <div class="subscribe-wechat-wrapper">
                    <input type="hidden" name="sceneStr" value="be87d7c37c294bfcbc159e4f4876cb95">
                    <input type="hidden" name="filter" value='{&quot;city&quot;:{&quot;desc&quot;:&quot;北京&quot;,&quot;value&quot;:&quot;101010100&quot;},&quot;filterMsg&quot;:&quot;【北京/golang/后端开发】&quot;,&quot;otherMsg&quot;:&quot;北京&quot;,&quot;position&quot;:{&quot;desc&quot;:&quot;后端开发&quot;,&quot;value&quot;:&quot;100199&quot;},&quot;query&quot;:&quot;golang&quot;,&quot;queryMsg&quot;:&quot;golang&quot;}'>
                    <a href="javascript:;" class="close" title="今天不再显示" ka="subscription_close"></a>
                    <dl>
                        <dt></dt>
                        <dd>微信扫一扫</dd>
                    </dl>
                    <i class="icon"></i>
                    <span class="title">新职位发布时通知我</span>
                    <p>订阅<span>【北京/golang/后端开发】</span>相关岗位，新岗位上线实时通知，求职快人一步</p>
                    <div class="sub-layer">
                        <div class="sub-container">
                            <p class="title">职位更新提醒</p >
                            <p class="info">新岗位已上线，找工作、快人一步</p >
                            <div class="content">
                                <div class="item">
                                    <div class="t">职位名称：</div>
                                    <div class="c">golang</div>
                                </div>
                                <div class="item">
                                    <div class="t">所在行业：</div>
                                    <div class="c">点击查看</div>
                                </div>
                                <div class="item">
                                    <div class="t">发布时间：</div>
                                    <div class="c">2021年9月12日</div>
                                </div>
                                <div class="item">
                                    <div class="t">备注：</div>
                                    <div class="c red">北京</div>
                                </div>
                            </div>
                            <div class="bottom"></div>
                        </div>
                    </div>
                </div>
            <div class="job-list">
                <div class="job-tab" style="display: none" data-filter="" data-keyword="golang" data-l3code="100199" data-rescount="271" data-page="1" data-source="1">
                    <a href="javascript:;" class="cur" ka="recommend-job-list">热门职位</a><a href="javascript:;" ka="new-job-list">最新职位</a>
                </div>

                    <div class="company-list">
                    </div>
                    <ul>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/79ef7ed97a0f9d5f1nFy0t6_FFRT.html" data-securityid="tCxp3WpRzaeY7-p1Vp4CIPYnX8GI--ZU5TkNoIQEE0yZPlh7N2Dc_LzzjIwy21By_ovv7ecqF-wOMc02RI4DGtkhEZOAyIo60ts6mBLRn8gWeCLpSg~~" data-jid="79ef7ed97a0f9d5f1nFy0t6_FFRT" data-itemid="1" data-lid="9d7sWA96TcY.search.1" data-jobid="158932463" data-index="0" ka="search_list_1" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/79ef7ed97a0f9d5f1nFy0t6_FFRT.html" title="Golang 研发" target="_blank" ka="search_list_jname_1" data-securityid="tCxp3WpRzaeY7-p1Vp4CIPYnX8GI--ZU5TkNoIQEE0yZPlh7N2Dc_LzzjIwy21By_ovv7ecqF-wOMc02RI4DGtkhEZOAyIo60ts6mBLRn8gWeCLpSg~~" data-jid="79ef7ed97a0f9d5f1nFy0t6_FFRT" data-itemid="1" data-lid="9d7sWA96TcY.search.1" data-jobid="158932463" data-index="0">Golang 研发</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·海淀区·上地</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">30-50K·15薪</span>
                                                    <p>3-5年<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>罗先生<em class="vline"></em>架构师</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_1" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=tCxp3WpRzaeY7-p1Vp4CIPYnX8GI--ZU5TkNoIQEE0yZPlh7N2Dc_LzzjIwy21By_ovv7ecqF-wOMc02RI4DGtkhEZOAyIo60ts6mBLRn8gWeCLpSg~~&jobId=79ef7ed97a0f9d5f1nFy0t6_FFRT&lid=9d7sWA96TcY.search.1" redirect-url="/web/geek/chat?id=420ca61bf5e15e661nZ53tq-FVI~">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/8548fadc0b5c265403V93t61EQ~~.html" title="滴滴招聘" ka="search_list_company_1_custompage" target="_blank">滴滴</a></h3>
                                                <p><a href="/i100019/" class="false-link" target="_blank" ka="search_list_company_industry_1_custompage" title="移动互联网行业招聘信息">移动互联网</a><em class="vline"></em>D轮及以上<em class="vline"></em>1000-9999人</p>
                                            </div>
                                            <a href="/gongsi/8548fadc0b5c265403V93t61EQ~~.html" ka="search_list_company_1_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/bar/20201205/25965fbe2c0440e760df2b6a6b81d7c2be1bd4a3bd2a63f070bdbdada9aad826.png?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">数据结构</span>
                                                    <span class="tag-item">架构师</span>
                                                    <span class="tag-item">Shell</span>
                                                    <span class="tag-item">微服务架构</span>
                                                    <span class="tag-item">分布式技术</span>
                                                    <span class="tag-item">Linux</span>
                                        </div>
                                        <div class="info-desc">五险一金，定期体检，补充医疗保险</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/62ca5a3a97897f171nBy39q-EFRX.html" data-securityid="V4mp-0TA1W6nj-N1vxdGY6fMej6iYuWv_YVZv_HjFTvtCJC9vVP36H05svOVTJAUGnWOOzYETIJwOSvhu1VqKJ5Cmb0c5691lsXSMXdsPGlN" data-jid="62ca5a3a97897f171nBy39q-EFRX" data-itemid="2" data-lid="9d7sWA96TcY.search.2" data-jobid="148473067" data-index="1" ka="search_list_2" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/62ca5a3a97897f171nBy39q-EFRX.html" title="安全开发 - Golang" target="_blank" ka="search_list_jname_2" data-securityid="V4mp-0TA1W6nj-N1vxdGY6fMej6iYuWv_YVZv_HjFTvtCJC9vVP36H05svOVTJAUGnWOOzYETIJwOSvhu1VqKJ5Cmb0c5691lsXSMXdsPGlN" data-jid="62ca5a3a97897f171nBy39q-EFRX" data-itemid="2" data-lid="9d7sWA96TcY.search.2" data-jobid="148473067" data-index="1">安全开发 - Golang</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·海淀区·颐和园</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">20-35K</span>
                                                    <p>5-10年<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>王先生<em class="vline"></em>研发副总裁</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_2" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=V4mp-0TA1W6nj-N1vxdGY6fMej6iYuWv_YVZv_HjFTvtCJC9vVP36H05svOVTJAUGnWOOzYETIJwOSvhu1VqKJ5Cmb0c5691lsXSMXdsPGlN&jobId=62ca5a3a97897f171nBy39q-EFRX&lid=9d7sWA96TcY.search.2" redirect-url="/web/geek/chat?id=2265156359484aa41Xd73d24">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/706cf439d50f694503V_2tU~.html" title="安天招聘" ka="search_list_company_2_custompage" target="_blank">安天</a></h3>
                                                <p><a href="/i100016/" class="false-link" target="_blank" ka="search_list_company_industry_2_custompage" title="信息安全行业招聘信息">信息安全</a><em class="vline"></em>C轮<em class="vline"></em>1000-9999人</p>
                                            </div>
                                            <a href="/gongsi/706cf439d50f694503V_2tU~.html" ka="search_list_company_2_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/chatphoto/20160405/ce5c90f6dfa276357c425da897c2fa3e0ac9c64003662f5fb5640db17e10ae92.jpg?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">Python</span>
                                                    <span class="tag-item">Java</span>
                                                    <span class="tag-item">中级软件工程师</span>
                                                    <span class="tag-item">网络安全</span>
                                                    <span class="tag-item">大数据</span>
                                                    <span class="tag-item">系统安全</span>
                                        </div>
                                        <div class="info-desc">定期体检，年终福利，补充医疗保险，免费班车，节日福利，外样补贴，年终奖，带薪年假，餐补，零食下午茶，五险一金</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/a11a928a811bad681nBz2NW_FlFU.html" data-securityid="udh2u8ljeVMRP-21rwjXSUVewn3tD5gQg2VD2tOEZBcP6_awJWc3K45LqKJ42oxx7JaSQHqWAPa0nAaGpG1_cyRXYdfbOmenxWHiNkMO5SYg-A~~" data-jid="a11a928a811bad681nBz2NW_FlFU" data-itemid="3" data-lid="9d7sWA96TcY.search.3" data-jobid="149382634" data-index="2" ka="search_list_3" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/a11a928a811bad681nBz2NW_FlFU.html" title="Golang 开发工程师" target="_blank" ka="search_list_jname_3" data-securityid="udh2u8ljeVMRP-21rwjXSUVewn3tD5gQg2VD2tOEZBcP6_awJWc3K45LqKJ42oxx7JaSQHqWAPa0nAaGpG1_cyRXYdfbOmenxWHiNkMO5SYg-A~~" data-jid="a11a928a811bad681nBz2NW_FlFU" data-itemid="3" data-lid="9d7sWA96TcY.search.3" data-jobid="149382634" data-index="2">Golang 开发工程师</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·朝阳区·酒仙桥</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">20-40K·14薪</span>
                                                    <p>1-3年<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>许女士<em class="vline"></em>招聘专员</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_3" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=udh2u8ljeVMRP-21rwjXSUVewn3tD5gQg2VD2tOEZBcP6_awJWc3K45LqKJ42oxx7JaSQHqWAPa0nAaGpG1_cyRXYdfbOmenxWHiNkMO5SYg-A~~&jobId=a11a928a811bad681nBz2NW_FlFU&lid=9d7sWA96TcY.search.3" redirect-url="/web/geek/chat?id=c9c1864a54fdfb1333B629-1FFs~">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/b758596cb99452101HZ-0g~~.html" title="美餐网招聘" ka="search_list_company_3_custompage" target="_blank">美餐网</a></h3>
                                                <p><a href="/i100020/" class="false-link" target="_blank" ka="search_list_company_industry_3_custompage" title="互联网行业招聘信息">互联网</a><em class="vline"></em>D轮及以上<em class="vline"></em>1000-9999人</p>
                                            </div>
                                            <a href="/gongsi/b758596cb99452101HZ-0g~~.html" ka="search_list_company_3_custompage_logo" target="_blank"><img class="company-logo" src="https://img2.bosszhipin.com/mcs/chatphoto/20151215/325f37c3a73265c95683c51ceb93badea6652a99e325a790f747a02d66801105.jpg?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">Python</span>
                                                    <span class="tag-item">Java</span>
                                                    <span class="tag-item">分布式技术</span>
                                        </div>
                                        <div class="info-desc">餐补，交通补助，年终奖，零食下午茶，通讯补贴，节日福利，员工旅游，带薪年假，五险一金，定期体检，股票期权</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/5c83fb58a841ba081nF83ti1F1dT.html" data-securityid="TJY28UOyCFXuc-a1eH7RsVjStML8B0Li9o6HkMaMiAEhxR2MrwAPOfpIT1Yo2NUGQ_0hhGdqee70kGyA04TQs6TjEVgFzxfbSJBbvMoQiDW0TKqf" data-jid="5c83fb58a841ba081nF83ti1F1dT" data-itemid="4" data-lid="9d7sWA96TcY.search.4" data-jobid="156558753" data-index="3" ka="search_list_4" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/5c83fb58a841ba081nF83ti1F1dT.html" title="golang 开发工程师" target="_blank" ka="search_list_jname_4" data-securityid="TJY28UOyCFXuc-a1eH7RsVjStML8B0Li9o6HkMaMiAEhxR2MrwAPOfpIT1Yo2NUGQ_0hhGdqee70kGyA04TQs6TjEVgFzxfbSJBbvMoQiDW0TKqf" data-jid="5c83fb58a841ba081nF83ti1F1dT" data-itemid="4" data-lid="9d7sWA96TcY.search.4" data-jobid="156558753" data-index="3">golang 开发工程师</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·朝阳区·三元桥</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">20-40K·14薪</span>
                                                    <p>3-5年<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>邱先生<em class="vline"></em>研究员</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_4" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=TJY28UOyCFXuc-a1eH7RsVjStML8B0Li9o6HkMaMiAEhxR2MrwAPOfpIT1Yo2NUGQ_0hhGdqee70kGyA04TQs6TjEVgFzxfbSJBbvMoQiDW0TKqf&jobId=5c83fb58a841ba081nF83ti1F1dT&lid=9d7sWA96TcY.search.4" redirect-url="/web/geek/chat?id=3bb807ab401225cc1nV63d6_FFtS">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/788e22a0c6957cbc1HN5298~.html" title="国美控股集团招聘" ka="search_list_company_4_custompage" target="_blank">国美控股集团</a></h3>
                                                <p><a href="/i100020/" class="false-link" target="_blank" ka="search_list_company_industry_4_custompage" title="互联网行业招聘信息">互联网</a><em class="vline"></em>不需要融资<em class="vline"></em>10000人以上</p>
                                            </div>
                                            <a href="/gongsi/788e22a0c6957cbc1HN5298~.html" ka="search_list_company_4_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/bar/20191129/bf060fc20f051ec80383f1d23aae96afbe1bd4a3bd2a63f070bdbdada9aad826.jpg?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">golang</span>
                                                    <span class="tag-item">后端开发</span>
                                                    <span class="tag-item">服务端开发</span>
                                                    <span class="tag-item">Gin</span>
                                                    <span class="tag-item">Beego</span>
                                                    <span class="tag-item">gRPC</span>
                                                    <span class="tag-item">go</span>
                                        </div>
                                        <div class="info-desc">年终奖，加班补助，节日福利，餐补，定期体检，带薪年假，通讯补贴，五险一金，员工旅游</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/79eda6123b8450651nd42du1FFdT.html" data-securityid="3edKL5fJd1I3n-X1QZcxl4cGKb3rQtjMn__riizhFmbckSNKT9fl45P5zlT2fHCpkx2WGjiwjx8AUk1oYXaqs9mmUmd_yRTXcFu17Maz0EuGoJE~" data-jid="79eda6123b8450651nd42du1FFdT" data-itemid="5" data-lid="9d7sWA96TcY.search.5" data-jobid="132268453" data-index="4" ka="search_list_5" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/79eda6123b8450651nd42du1FFdT.html" title="golang 研发工程师" target="_blank" ka="search_list_jname_5" data-securityid="3edKL5fJd1I3n-X1QZcxl4cGKb3rQtjMn__riizhFmbckSNKT9fl45P5zlT2fHCpkx2WGjiwjx8AUk1oYXaqs9mmUmd_yRTXcFu17Maz0EuGoJE~" data-jid="79eda6123b8450651nd42du1FFdT" data-itemid="5" data-lid="9d7sWA96TcY.search.5" data-jobid="132268453" data-index="4">golang 研发工程师</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">25-40K·14薪</span>
                                                    <p>3-5年<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>刘女士<em class="vline"></em>HR</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_5" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=3edKL5fJd1I3n-X1QZcxl4cGKb3rQtjMn__riizhFmbckSNKT9fl45P5zlT2fHCpkx2WGjiwjx8AUk1oYXaqs9mmUmd_yRTXcFu17Maz0EuGoJE~&jobId=79eda6123b8450651nd42du1FFdT&lid=9d7sWA96TcY.search.5" redirect-url="/web/geek/chat?id=2aa2f8dc01630c5b1HJ92Ni9Eg~~">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/3cd0c8b96fd0bd001nJ53dy5.html" title="长亭科技招聘" ka="search_list_company_5_custompage" target="_blank">长亭科技</a></h3>
                                                <p><a href="/i100016/" class="false-link" target="_blank" ka="search_list_company_industry_5_custompage" title="信息安全行业招聘信息">信息安全</a><em class="vline"></em>不需要融资<em class="vline"></em>500-999人</p>
                                            </div>
                                            <a href="/gongsi/3cd0c8b96fd0bd001nJ53dy5.html" ka="search_list_company_5_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/chatphoto/20170920/1ee980ff4946820864c5875d4d17756a904a2a9f4c059ef37822602fbb68a580.jpg?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">Shell</span>
                                                    <span class="tag-item">分布式技术</span>
                                                    <span class="tag-item">容器技术</span>
                                                    <span class="tag-item">Linux</span>
                                                    <span class="tag-item">架构师</span>
                                        </div>
                                        <div class="info-desc">零食下午茶，餐补，年终奖，包吃，带薪年假，节日福利，定期体检，补充医疗保险，交通补助，五险一金，通讯补贴，员工旅游</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/b94e92b49a466dc81nF53tu_GFpQ.html" data-securityid="4Vil0xgGG72Kf-T1c428CDE6cvwkmPAJhrbYchIj75PPUifIhVUkHxt3jH1KqdyJWhKHCGvOI7pDOyNeHbWEcVvbphNy8KTIgrmTVcmr-Jaw" data-jid="b94e92b49a466dc81nF53tu_GFpQ" data-itemid="6" data-lid="9d7sWA96TcY.search.6" data-jobid="153562880" data-index="5" ka="search_list_6" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/b94e92b49a466dc81nF53tu_GFpQ.html" title="高级 Golang 工程师" target="_blank" ka="search_list_jname_6" data-securityid="4Vil0xgGG72Kf-T1c428CDE6cvwkmPAJhrbYchIj75PPUifIhVUkHxt3jH1KqdyJWhKHCGvOI7pDOyNeHbWEcVvbphNy8KTIgrmTVcmr-Jaw" data-jid="b94e92b49a466dc81nF53tu_GFpQ" data-itemid="6" data-lid="9d7sWA96TcY.search.6" data-jobid="153562880" data-index="5">高级 Golang 工程师</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·海淀区·上地</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">20-35K</span>
                                                    <p>3-5年<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>李先生<em class="vline"></em>研发总监</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_6" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=4Vil0xgGG72Kf-T1c428CDE6cvwkmPAJhrbYchIj75PPUifIhVUkHxt3jH1KqdyJWhKHCGvOI7pDOyNeHbWEcVvbphNy8KTIgrmTVcmr-Jaw&jobId=b94e92b49a466dc81nF53tu_GFpQ&lid=9d7sWA96TcY.search.6" redirect-url="/web/geek/chat?id=486fa44d0676b0050nN73Ny6Eg~~">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/8de1f130f2c1340d0XZ_2g~~.html" title="青藤云安全招聘" ka="search_list_company_6_custompage" target="_blank">青藤云安全</a></h3>
                                                <p><a href="/i100016/" class="false-link" target="_blank" ka="search_list_company_industry_6_custompage" title="信息安全行业招聘信息">信息安全</a><em class="vline"></em>C轮<em class="vline"></em>500-999人</p>
                                            </div>
                                            <a href="/gongsi/8de1f130f2c1340d0XZ_2g~~.html" ka="search_list_company_6_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/chatphoto/20190415/c06d78678c87b4d75f8c6dc0b0f3a5de660212d513ed3f17d71f71b5bf054f9b.jpg?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">分布式技术</span>
                                                    <span class="tag-item">容器技术</span>
                                                    <span class="tag-item">Linux</span>
                                                    <span class="tag-item">Shell</span>
                                                    <span class="tag-item">网络编程</span>
                                                    <span class="tag-item">微服务架构</span>
                                                    <span class="tag-item">后端开发</span>
                                                    <span class="tag-item">高级软件工程师</span>
                                        </div>
                                        <div class="info-desc">定期体检，带薪年假，年终奖，餐补，股票期权，五险一金，零食下午茶，节日福利</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/d833fdcce2de63030XJ70tW_FVo~.html" data-securityid="mvy2w6YyFrHwB-f1nFbObfzJOzlq6N60Eaeg6n5tRMRBOuzvu0GBr6etzgzL3EPXVzPGiaGxuwUDVMLIN6pE-mT4poKH-I1KOX3O3CWe2ahj7GG7BA~~" data-jid="d833fdcce2de63030XJ70tW_FVo~" data-itemid="7" data-lid="9d7sWA96TcY.search.7" data-jobid="66198258" data-index="6" ka="search_list_7" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/d833fdcce2de63030XJ70tW_FVo~.html" title="Golang 后端研发工程师" target="_blank" ka="search_list_jname_7" data-securityid="mvy2w6YyFrHwB-f1nFbObfzJOzlq6N60Eaeg6n5tRMRBOuzvu0GBr6etzgzL3EPXVzPGiaGxuwUDVMLIN6pE-mT4poKH-I1KOX3O3CWe2ahj7GG7BA~~" data-jid="d833fdcce2de63030XJ70tW_FVo~" data-itemid="7" data-lid="9d7sWA96TcY.search.7" data-jobid="66198258" data-index="6">Golang 后端研发工程师</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·东城区·崇文门</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">20-35K</span>
                                                    <p>1-3年<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>特先生<em class="vline"></em>工程VP</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_7" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=mvy2w6YyFrHwB-f1nFbObfzJOzlq6N60Eaeg6n5tRMRBOuzvu0GBr6etzgzL3EPXVzPGiaGxuwUDVMLIN6pE-mT4poKH-I1KOX3O3CWe2ahj7GG7BA~~&jobId=d833fdcce2de63030XJ70tW_FVo~&lid=9d7sWA96TcY.search.7" redirect-url="/web/geek/chat?id=1bcb76752631bd9d1HN-2ty1EFY~">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/db6e5eb843cb8a921nFz0t25EVI~.html" title="科亚医疗科技股份...招聘" ka="search_list_company_7_custompage" target="_blank">科亚医疗科技股份...</a></h3>
                                                <p><a href="/i100403/" class="false-link" target="_blank" ka="search_list_company_industry_7_custompage" title="医疗设备/器械行业招聘信息">医疗设备/器械</a><em class="vline"></em>D轮及以上<em class="vline"></em>500-999人</p>
                                            </div>
                                            <a href="/gongsi/db6e5eb843cb8a921nFz0t25EVI~.html" ka="search_list_company_7_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/upload/com/workfeel/20210821/7bf6f160950405e98bdaa7462a55e66a9fa0d8d8df4936bf08aedf49ff775c085bc0b36ed031672a.jpg?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">Go</span>
                                                    <span class="tag-item">容器技术</span>
                                                    <span class="tag-item">后端开发</span>
                                                    <span class="tag-item">Linux</span>
                                                    <span class="tag-item">GIT</span>
                                                    <span class="tag-item">分布式技术</span>
                                                    <span class="tag-item">云计算</span>
                                                    <span class="tag-item">大数据</span>
                                        </div>
                                        <div class="info-desc">通讯补贴，年终奖，五险一金，餐补，零食下午茶</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/0247c04716c16c301nJ439y7E1dW.html" data-securityid="ToqL7tkVpoC2p-41UWrqH2FVsXItyxVccOs0gQloG9Pwbcd0w1eGK9mdpMK7AHinqEQuDByjhJYfQ8ZNlJAM2xM2Bk7J1rYtYx33J7oHyoWxpgQ~" data-jid="0247c04716c16c301nJ439y7E1dW" data-itemid="8" data-lid="9d7sWA96TcY.search.8" data-jobid="162416356" data-index="7" ka="search_list_8" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/0247c04716c16c301nJ439y7E1dW.html" title="Golang 研发工程师" target="_blank" ka="search_list_jname_8" data-securityid="ToqL7tkVpoC2p-41UWrqH2FVsXItyxVccOs0gQloG9Pwbcd0w1eGK9mdpMK7AHinqEQuDByjhJYfQ8ZNlJAM2xM2Bk7J1rYtYx33J7oHyoWxpgQ~" data-jid="0247c04716c16c301nJ439y7E1dW" data-itemid="8" data-lid="9d7sWA96TcY.search.8" data-jobid="162416356" data-index="7">Golang 研发工程师</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·海淀区·清河</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">25-50K</span>
                                                    <p>3-5年<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>李女士<em class="vline"></em>小米 HR</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_8" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=ToqL7tkVpoC2p-41UWrqH2FVsXItyxVccOs0gQloG9Pwbcd0w1eGK9mdpMK7AHinqEQuDByjhJYfQ8ZNlJAM2xM2Bk7J1rYtYx33J7oHyoWxpgQ~&jobId=0247c04716c16c301nJ439y7E1dW&lid=9d7sWA96TcY.search.8" redirect-url="/web/geek/chat?id=86aebc17c5144f461Hd_3Ni0GVc~">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/6f1aa1d6b1d033ad33B43N0~.html" title="小米招聘" ka="search_list_company_8_custompage" target="_blank">小米</a></h3>
                                                <p><a href="/i100020/" class="false-link" target="_blank" ka="search_list_company_industry_8_custompage" title="互联网行业招聘信息">互联网</a><em class="vline"></em>已上市<em class="vline"></em>10000人以上</p>
                                            </div>
                                            <a href="/gongsi/6f1aa1d6b1d033ad33B43N0~.html" ka="search_list_company_8_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/chatphoto/20210331/53a0321d3b2965103208996a0b0cf109eae606b60d295d07a1b828b97fa2d778_s.jpg?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">数据结构</span>
                                                    <span class="tag-item">Linux</span>
                                                    <span class="tag-item">golang</span>
                                                    <span class="tag-item">云平台</span>
                                                    <span class="tag-item">运维开发</span>
                                        </div>
                                        <div class="info-desc">全勤奖，股票期权，带薪年假，五险一金，12%公积金，年终奖，补充医疗保险，餐补，节日福利，定期体检</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/2182f9413e3fd13c1nBz39q7FFZW.html" data-securityid="pZFgw9HKZj7GB-v1XlLgZEbHutbVq-NjVt1xi0wGgFVhIqfLfdXPEEIKPbyi8HcRP-3mm5L2Uc3Y3zuQ8BoDVYWp6TZ-bgM205_xd_r1pJqWxQ~~" data-jid="2182f9413e3fd13c1nBz39q7FFZW" data-itemid="9" data-lid="9d7sWA96TcY.search.9" data-jobid="149476446" data-index="8" ka="search_list_9" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/2182f9413e3fd13c1nBz39q7FFZW.html" title="Golang 高级研发工程师" target="_blank" ka="search_list_jname_9" data-securityid="pZFgw9HKZj7GB-v1XlLgZEbHutbVq-NjVt1xi0wGgFVhIqfLfdXPEEIKPbyi8HcRP-3mm5L2Uc3Y3zuQ8BoDVYWp6TZ-bgM205_xd_r1pJqWxQ~~" data-jid="2182f9413e3fd13c1nBz39q7FFZW" data-itemid="9" data-lid="9d7sWA96TcY.search.9" data-jobid="149476446" data-index="8">Golang 高级研发工程师</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·海淀区·学院路</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">25-50K·15薪</span>
                                                    <p>经验不限<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>唱先生<em class="vline"></em>内容中心.移动端负责人</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_9" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=pZFgw9HKZj7GB-v1XlLgZEbHutbVq-NjVt1xi0wGgFVhIqfLfdXPEEIKPbyi8HcRP-3mm5L2Uc3Y3zuQ8BoDVYWp6TZ-bgM205_xd_r1pJqWxQ~~&jobId=2182f9413e3fd13c1nBz39q7FFZW&lid=9d7sWA96TcY.search.9" redirect-url="/web/geek/chat?id=58edbf871f1622251HVz3du4">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/8372b71df55405071nJ52Nu9.html" title="知乎招聘" ka="search_list_company_9_custompage" target="_blank">知乎</a></h3>
                                                <p><a href="/i100019/" class="false-link" target="_blank" ka="search_list_company_industry_9_custompage" title="移动互联网行业招聘信息">移动互联网</a><em class="vline"></em>已上市<em class="vline"></em>1000-9999人</p>
                                            </div>
                                            <a href="/gongsi/8372b71df55405071nJ52Nu9.html" ka="search_list_company_9_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/bar/20201205/2437636f1a3e49d311fa6d19e02d0f5ebe1bd4a3bd2a63f070bdbdada9aad826.png?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">后端开发</span>
                                                    <span class="tag-item">服务端开发</span>
                                                    <span class="tag-item">分布式技术</span>
                                                    <span class="tag-item">数据结构</span>
                                                    <span class="tag-item">网络协议</span>
                                                    <span class="tag-item">高级软件工程师</span>
                                        </div>
                                        <div class="info-desc">节日福利，定期体检，零食下午茶，员工旅游，年终奖，补充医疗保险，股票期权，五险一金，包吃，带薪年假</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/6e8e726f37aa9e971nd42tW1EVFT.html" data-securityid="RssDhxsf4xsZ4-l1Sz1FMmnp1RvT-jyTjysEm5MXJsuWuXSIRJ_vBEHw3ChQqTBmSPfEaEIiic3Jz4-STqYzsyGDIQwlLFVsJeKIvF3Ru6A8inI~" data-jid="6e8e726f37aa9e971nd42tW1EVFT" data-itemid="10" data-lid="9d7sWA96TcY.search.10" data-jobid="132188133" data-index="9" ka="search_list_10" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/6e8e726f37aa9e971nd42tW1EVFT.html" title="Golang 开发工程师" target="_blank" ka="search_list_jname_10" data-securityid="RssDhxsf4xsZ4-l1Sz1FMmnp1RvT-jyTjysEm5MXJsuWuXSIRJ_vBEHw3ChQqTBmSPfEaEIiic3Jz4-STqYzsyGDIQwlLFVsJeKIvF3Ru6A8inI~" data-jid="6e8e726f37aa9e971nd42tW1EVFT" data-itemid="10" data-lid="9d7sWA96TcY.search.10" data-jobid="132188133" data-index="9">Golang 开发工程师</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·海淀区·清河</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">15-30K·13薪</span>
                                                    <p>3-5年<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>付先生<em class="vline"></em>架构师</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_10" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=RssDhxsf4xsZ4-l1Sz1FMmnp1RvT-jyTjysEm5MXJsuWuXSIRJ_vBEHw3ChQqTBmSPfEaEIiic3Jz4-STqYzsyGDIQwlLFVsJeKIvF3Ru6A8inI~&jobId=6e8e726f37aa9e971nd42tW1EVFT&lid=9d7sWA96TcY.search.10" redirect-url="/web/geek/chat?id=7b15f77001b58f561nR62dq4EVdY">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/b4f11fb7a32c37761nB62Q~~.html" title="一亩田招聘" ka="search_list_company_10_custompage" target="_blank">一亩田</a></h3>
                                                <p><a href="/i100008/" class="false-link" target="_blank" ka="search_list_company_industry_10_custompage" title="O2O行业招聘信息">O2O</a><em class="vline"></em>C轮<em class="vline"></em>500-999人</p>
                                            </div>
                                            <a href="/gongsi/b4f11fb7a32c37761nB62Q~~.html" ka="search_list_company_10_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/chatphoto/20190216/2249e96fd4d957df2f1b529ec8b2bceccfcd208495d565ef66e7dff9f98764da.jpg?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">PHP</span>
                                                    <span class="tag-item">数据结构</span>
                                                    <span class="tag-item">Linux</span>
                                                    <span class="tag-item">用户中心</span>
                                                    <span class="tag-item">Golang</span>
                                                    <span class="tag-item">IAM</span>
                                                    <span class="tag-item">领域架构</span>
                                        </div>
                                        <div class="info-desc">带薪年假，全勤奖，年终奖，定期体检，交通补助，零食下午茶，住房补贴，节日福利，股票期权，补充医疗保险，五险一金</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/c4edc2ab8648f33b1nJ92d-9EVpU.html" data-securityid="J1nZ7XTru9CK--K1NXxM_0o2_pjyhT0TRdnGcGp3cgIfdSiRwZ878TrjffxjG6M5H2kdoAZ-gw1b0R6GFpTdAIpAv5LTGpLYGh4swSnO8o3XY6U~" data-jid="c4edc2ab8648f33b1nJ92d-9EVpU" data-itemid="11" data-lid="9d7sWA96TcY.search.11" data-jobid="167220184" data-index="10" ka="search_list_11" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/c4edc2ab8648f33b1nJ92d-9EVpU.html" title="Golang 后端工程师" target="_blank" ka="search_list_jname_11" data-securityid="J1nZ7XTru9CK--K1NXxM_0o2_pjyhT0TRdnGcGp3cgIfdSiRwZ878TrjffxjG6M5H2kdoAZ-gw1b0R6GFpTdAIpAv5LTGpLYGh4swSnO8o3XY6U~" data-jid="c4edc2ab8648f33b1nJ92d-9EVpU" data-itemid="11" data-lid="9d7sWA96TcY.search.11" data-jobid="167220184" data-index="10">Golang 后端工程师</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·东城区·东单</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">20-40K·14薪</span>
                                                    <p>1-3年<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>翁女士<em class="vline"></em>HRBP</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_11" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=J1nZ7XTru9CK--K1NXxM_0o2_pjyhT0TRdnGcGp3cgIfdSiRwZ878TrjffxjG6M5H2kdoAZ-gw1b0R6GFpTdAIpAv5LTGpLYGh4swSnO8o3XY6U~&jobId=c4edc2ab8648f33b1nJ92d-9EVpU&lid=9d7sWA96TcY.search.11" redirect-url="/web/geek/chat?id=4c0b3029ae6a0e740HZ-0ti9EA~~">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/7ec358e71cf9cc4d1X1z3Ny_.html" title="Opera招聘" ka="search_list_company_11_custompage" target="_blank">Opera</a></h3>
                                                <p><a href="/i100019/" class="false-link" target="_blank" ka="search_list_company_industry_11_custompage" title="移动互联网行业招聘信息">移动互联网</a><em class="vline"></em>已上市<em class="vline"></em>500-999人</p>
                                            </div>
                                            <a href="/gongsi/7ec358e71cf9cc4d1X1z3Ny_.html" ka="search_list_company_11_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/bar/20200618/647f778610f5b88145fce8baa3db7c52be1bd4a3bd2a63f070bdbdada9aad826.png?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">Python</span>
                                                    <span class="tag-item">Java</span>
                                                    <span class="tag-item">容器技术</span>
                                                    <span class="tag-item">Beego</span>
                                                    <span class="tag-item">Gin</span>
                                        </div>
                                        <div class="info-desc">年终奖，带薪年假，补充医疗保险，节日福利，零食下午茶，定期体检，餐补，包吃，员工旅游，五险一金</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/31a10602f8d600fc1nZ72dm_EVFY.html" data-securityid="tXEfqMXBBBKrL-V16-h1RIxHlwSVrxW4sd_Dgtdgd56o9S4fRqqHsqP7VXk6uI-0ZTbzOx-7pa7UrfcA9JKx3CMe77Bal3MPR68CFOKC4yeAnZw~" data-jid="31a10602f8d600fc1nZ72dm_EVFY" data-itemid="12" data-lid="9d7sWA96TcY.search.12" data-jobid="121242138" data-index="11" ka="search_list_12" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/31a10602f8d600fc1nZ72dm_EVFY.html" title="Senior Golang Engineer" target="_blank" ka="search_list_jname_12" data-securityid="tXEfqMXBBBKrL-V16-h1RIxHlwSVrxW4sd_Dgtdgd56o9S4fRqqHsqP7VXk6uI-0ZTbzOx-7pa7UrfcA9JKx3CMe77Bal3MPR68CFOKC4yeAnZw~" data-jid="31a10602f8d600fc1nZ72dm_EVFY" data-itemid="12" data-lid="9d7sWA96TcY.search.12" data-jobid="121242138" data-index="11">Senior Golang Engineer</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">25-35K·14薪</span>
                                                    <p>3-5年<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>邓先生<em class="vline"></em>HR</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_12" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=tXEfqMXBBBKrL-V16-h1RIxHlwSVrxW4sd_Dgtdgd56o9S4fRqqHsqP7VXk6uI-0ZTbzOx-7pa7UrfcA9JKx3CMe77Bal3MPR68CFOKC4yeAnZw~&jobId=31a10602f8d600fc1nZ72dm_EVFY&lid=9d7sWA96TcY.search.12" redirect-url="/web/geek/chat?id=84c32d15df1073b81nZ92tu-EFs~">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/eb36cd62c98680371nN82NQ~.html" title="FreeWheel招聘" ka="search_list_company_12_custompage" target="_blank">FreeWheel</a></h3>
                                                <p><a href="/i100020/" class="false-link" target="_blank" ka="search_list_company_industry_12_custompage" title="互联网行业招聘信息">互联网</a><em class="vline"></em>已上市<em class="vline"></em>500-999人</p>
                                            </div>
                                            <a href="/gongsi/eb36cd62c98680371nN82NQ~.html" ka="search_list_company_12_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/chatphoto/20180117/bfa0c218d81af2812f2d4e85b1d787b7cfcd208495d565ef66e7dff9f98764da.jpg?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">Golang</span>
                                        </div>
                                        <div class="info-desc">节日福利，免费班车，定期体检，五险一金，员工旅游，加班补助，通讯补贴，股票期权，包吃，零食下午茶，餐补，交通补助，年终奖，补充医疗保险，带薪年假</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/03e479f5781f4cdf1nZ509y_FFVR.html" data-securityid="z58k4j5SHADK0-K1Si1BvKPDVt0yV_v_IPFh0XBHFfPcpXCMVbOjXkfliXMWrcpGzfaWE4E3nSPCijEz-vZwMLaL4j5WdFi4kE3bON8w4f2ZCJM~" data-jid="03e479f5781f4cdf1nZ509y_FFVR" data-itemid="13" data-lid="9d7sWA96TcY.search.13" data-jobid="123812471" data-index="12" ka="search_list_13" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/03e479f5781f4cdf1nZ509y_FFVR.html" title="高级 golang 开发工程师" target="_blank" ka="search_list_jname_13" data-securityid="z58k4j5SHADK0-K1Si1BvKPDVt0yV_v_IPFh0XBHFfPcpXCMVbOjXkfliXMWrcpGzfaWE4E3nSPCijEz-vZwMLaL4j5WdFi4kE3bON8w4f2ZCJM~" data-jid="03e479f5781f4cdf1nZ509y_FFVR" data-itemid="13" data-lid="9d7sWA96TcY.search.13" data-jobid="123812471" data-index="12">高级 golang 开发工程师</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·海淀区·万寿路</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">25-50K·14薪</span>
                                                    <p>5-10年<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>郑先生<em class="vline"></em>DevOps架构师</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_13" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=z58k4j5SHADK0-K1Si1BvKPDVt0yV_v_IPFh0XBHFfPcpXCMVbOjXkfliXMWrcpGzfaWE4E3nSPCijEz-vZwMLaL4j5WdFi4kE3bON8w4f2ZCJM~&jobId=03e479f5781f4cdf1nZ509y_FFVR&lid=9d7sWA96TcY.search.13" redirect-url="/web/geek/chat?id=477dc3c4687c5cff1nB92d2_FA~~">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/a6e0e2078e7d54f01nB70ty6.html" title="锐捷网络招聘" ka="search_list_company_13_custompage" target="_blank">锐捷网络</a></h3>
                                                <p><a href="/i100024/" class="false-link" target="_blank" ka="search_list_company_industry_13_custompage" title="通信/网络设备行业招聘信息">通信/网络设备</a><em class="vline"></em>已上市<em class="vline"></em>1000-9999人</p>
                                            </div>
                                            <a href="/gongsi/a6e0e2078e7d54f01nB70ty6.html" ka="search_list_company_13_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/chatphoto/20160817/8d246f951d426de875a95246f21e89490045bc52c878d1a70ccd15e2d1e5a348.jpg?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">分布式技术</span>
                                                    <span class="tag-item">数据结构</span>
                                                    <span class="tag-item">容器技术</span>
                                                    <span class="tag-item">消息队列</span>
                                                    <span class="tag-item">Linux</span>
                                                    <span class="tag-item">Unix</span>
                                                    <span class="tag-item">架构师</span>
                                                    <span class="tag-item">AI</span>
                                        </div>
                                        <div class="info-desc">餐补，补充医疗保险，交通补助，通讯补贴，带薪年假，年终奖，员工旅游，加班补助，五险一金，免费班车，爱心基金，节日福利，零食下午茶，包吃，定期体检</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/be581693e5e026f41nJy29q-FVdR.html" data-securityid="qAN9f9dyKrc3q-S1L8-AbQdfC49udjX_fSsMRNfF5vBzSZ9647HctP4b92BYSqRJpqIuWg7zKLuoUVISNXkj23T72F5-_Y4KOGAZEjw56oVcbdJk1A~~" data-jid="be581693e5e026f41nJy29q-FVdR" data-itemid="14" data-lid="9d7sWA96TcY.search.14" data-jobid="168073551" data-index="13" ka="search_list_14" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/be581693e5e026f41nJy29q-FVdR.html" title="Senior Golang API Developer" target="_blank" ka="search_list_jname_14" data-securityid="qAN9f9dyKrc3q-S1L8-AbQdfC49udjX_fSsMRNfF5vBzSZ9647HctP4b92BYSqRJpqIuWg7zKLuoUVISNXkj23T72F5-_Y4KOGAZEjw56oVcbdJk1A~~" data-jid="be581693e5e026f41nJy29q-FVdR" data-itemid="14" data-lid="9d7sWA96TcY.search.14" data-jobid="168073551" data-index="13">Senior Golang API Developer</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·朝阳区·三里屯</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">35-65K·13薪</span>
                                                    <p>5-10年<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>周女士<em class="vline"></em>TA Manager</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_14" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=qAN9f9dyKrc3q-S1L8-AbQdfC49udjX_fSsMRNfF5vBzSZ9647HctP4b92BYSqRJpqIuWg7zKLuoUVISNXkj23T72F5-_Y4KOGAZEjw56oVcbdJk1A~~&jobId=be581693e5e026f41nJy29q-FVdR&lid=9d7sWA96TcY.search.14" redirect-url="/web/geek/chat?id=ef57071ec659319a1nR_09i6GVpQ">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/d5c3f9e25ae817001nJ93Nm4.html" title="伽蓝集团招聘" ka="search_list_company_14_custompage" target="_blank">伽蓝集团</a></h3>
                                                <p><a href="/i100001/" class="false-link" target="_blank" ka="search_list_company_industry_14_custompage" title="电子商务行业招聘信息">电子商务</a><em class="vline"></em>不需要融资<em class="vline"></em>1000-9999人</p>
                                            </div>
                                            <a href="/gongsi/d5c3f9e25ae817001nJ93Nm4.html" ka="search_list_company_14_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/chatphoto/20160727/79057cc80cb8042ebec79df9e8861f9a365d64227d57fb638341c71cac540676.jpg?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">gin</span>
                                                    <span class="tag-item">golang</span>
                                        </div>
                                        <div class="info-desc">全勤奖，年终奖，补充医疗保险，免费班车，加班补助，五险一金，带薪年假，节日福利，交通补助，餐补，定期体检</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/ba0e743842c391951nx409q_FVNZ.html" data-securityid="YYZX1pzu1uNhg-H1PaQ8pab5NdFgg2yQ9KrQA68RmXChKfyXO9XL3AuPaoMmq3MRRKvzV8WWN9Lb8ZkPDd9vwUniEv_edtXNndbwndmNxCrpY8P7" data-jid="ba0e743842c391951nx409q_FVNZ" data-itemid="15" data-lid="9d7sWA96TcY.search.15" data-jobid="182872519" data-index="14" ka="search_list_15" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/ba0e743842c391951nx409q_FVNZ.html" title="Golang" target="_blank" ka="search_list_jname_15" data-securityid="YYZX1pzu1uNhg-H1PaQ8pab5NdFgg2yQ9KrQA68RmXChKfyXO9XL3AuPaoMmq3MRRKvzV8WWN9Lb8ZkPDd9vwUniEv_edtXNndbwndmNxCrpY8P7" data-jid="ba0e743842c391951nx409q_FVNZ" data-itemid="15" data-lid="9d7sWA96TcY.search.15" data-jobid="182872519" data-index="14">Golang</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·海淀区·皂君庙</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">20-40K·15薪</span>
                                                    <p>经验不限<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>吴先生<em class="vline"></em>西瓜服务端开发工程师</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_15" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=YYZX1pzu1uNhg-H1PaQ8pab5NdFgg2yQ9KrQA68RmXChKfyXO9XL3AuPaoMmq3MRRKvzV8WWN9Lb8ZkPDd9vwUniEv_edtXNndbwndmNxCrpY8P7&jobId=ba0e743842c391951nx409q_FVNZ&lid=9d7sWA96TcY.search.15" redirect-url="/web/geek/chat?id=7868b8cc46c08edd0HJ53dS8Els~">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/7d01ce6cfe2022030HJ409y-.html" title="北京字节跳动招聘" ka="search_list_company_15_custompage" target="_blank">北京字节跳动</a></h3>
                                                <p><a href="/i100019/" class="false-link" target="_blank" ka="search_list_company_industry_15_custompage" title="移动互联网行业招聘信息">移动互联网</a><em class="vline"></em>D轮及以上<em class="vline"></em>10000人以上</p>
                                            </div>
                                            <a href="/gongsi/7d01ce6cfe2022030HJ409y-.html" ka="search_list_company_15_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/chatphoto/20201123/424d60a634b16f6d498bb81a1ecc4fb882e5997348729a6cfe816ad90c892e55_s.jpg?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">服务端开发</span>
                                                    <span class="tag-item">golang</span>
                                        </div>
                                        <div class="info-desc">六险一金，餐补，嘀嘀打车，免费健身房，免费晚餐，零食下午茶，员工旅游，交通补助，定期体检，住房补贴，带薪年假，年终奖，五险一金，补充医疗保险，包吃，节日福利，六险一金，股票期权</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/44cdb794ea9589711nd72d-1FVBW.html" data-securityid="5LJHgicFal5gD-V1V_Pdk-Wu4K5wVvPN_nqpH9V0kbmT9_3ospJ80IcvInvg7knIjX9xZ58khvCQm72l81g6gYi0426-3OIlm4pTD8h-KiS5Mn6c-A~~" data-jid="44cdb794ea9589711nd72d-1FVBW" data-itemid="16" data-lid="9d7sWA96TcY.search.16" data-jobid="131228526" data-index="15" ka="search_list_16" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/44cdb794ea9589711nd72d-1FVBW.html" title="Golang" target="_blank" ka="search_list_jname_16" data-securityid="5LJHgicFal5gD-V1V_Pdk-Wu4K5wVvPN_nqpH9V0kbmT9_3ospJ80IcvInvg7knIjX9xZ58khvCQm72l81g6gYi0426-3OIlm4pTD8h-KiS5Mn6c-A~~" data-jid="44cdb794ea9589711nd72d-1FVBW" data-itemid="16" data-lid="9d7sWA96TcY.search.16" data-jobid="131228526" data-index="15">Golang</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·西城区·宣武门</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">15-30K·15薪</span>
                                                    <p>经验不限<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>康先生<em class="vline"></em>云计算工程师</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_16" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=5LJHgicFal5gD-V1V_Pdk-Wu4K5wVvPN_nqpH9V0kbmT9_3ospJ80IcvInvg7knIjX9xZ58khvCQm72l81g6gYi0426-3OIlm4pTD8h-KiS5Mn6c-A~~&jobId=44cdb794ea9589711nd72d-1FVBW&lid=9d7sWA96TcY.search.16" redirect-url="/web/geek/chat?id=17f5d251365335ff0nN63NW0E1s~">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/1bedf6efac8cbf6b0Hd62NW6EQ~~.html" title="中移雄安招聘" ka="search_list_company_16_custompage" target="_blank">中移雄安</a></h3>
                                                <p><a href="/i100019/" class="false-link" target="_blank" ka="search_list_company_industry_16_custompage" title="移动互联网行业招聘信息">移动互联网</a><em class="vline"></em>不需要融资<em class="vline"></em>1000-9999人</p>
                                            </div>
                                            <a href="/gongsi/1bedf6efac8cbf6b0Hd62NW6EQ~~.html" ka="search_list_company_16_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/upload/com/logo/20200304/c4c208e6494d5aeee7da7700abb27c584f83d7cfc75c309a7f2be224cb4ff19f.jpg?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">容器技术</span>
                                                    <span class="tag-item">Linux</span>
                                                    <span class="tag-item">云计算</span>
                                                    <span class="tag-item">kubernetes</span>
                                                    <span class="tag-item">golang</span>
                                        </div>
                                        <div class="info-desc">六险二金，年终奖，包吃，住房补贴，定期体检，五险一金，节日福利，通讯补贴，不加班，餐补，带薪年假，零食下午茶，补充医疗保险</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/e7a655fafd6ce7df1nJ-2tu9ElNW.html" data-securityid="C67gTxumx2RPN-e1-5UiAJ76h5uzXE6Pn_a9MNXC1BXbNsMIJWlR-H2ift9PpatXQHLHYBZYiYSmsZY2bTQkJxKpkkF28kR7WddpIeVXFG3I" data-jid="e7a655fafd6ce7df1nJ-2tu9ElNW" data-itemid="17" data-lid="9d7sWA96TcY.search.17" data-jobid="164160216" data-index="16" ka="search_list_17" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/e7a655fafd6ce7df1nJ-2tu9ElNW.html" title="Golang" target="_blank" ka="search_list_jname_17" data-securityid="C67gTxumx2RPN-e1-5UiAJ76h5uzXE6Pn_a9MNXC1BXbNsMIJWlR-H2ift9PpatXQHLHYBZYiYSmsZY2bTQkJxKpkkF28kR7WddpIeVXFG3I" data-jid="e7a655fafd6ce7df1nJ-2tu9ElNW" data-itemid="17" data-lid="9d7sWA96TcY.search.17" data-jobid="164160216" data-index="16">Golang</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·海淀区·西北旺</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">25-40K·14薪</span>
                                                    <p>1-3年<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>王先生<em class="vline"></em>后端开发</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_17" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=C67gTxumx2RPN-e1-5UiAJ76h5uzXE6Pn_a9MNXC1BXbNsMIJWlR-H2ift9PpatXQHLHYBZYiYSmsZY2bTQkJxKpkkF28kR7WddpIeVXFG3I&jobId=e7a655fafd6ce7df1nJ-2tu9ElNW&lid=9d7sWA96TcY.search.17" redirect-url="/web/geek/chat?id=931e23837a2e3cc13nx929i7">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/f3e62b9ed1bfa70b1nNy09w~.html" title="微博招聘" ka="search_list_company_17_custompage" target="_blank">微博</a></h3>
                                                <p><a href="/i100020/" class="false-link" target="_blank" ka="search_list_company_industry_17_custompage" title="互联网行业招聘信息">互联网</a><em class="vline"></em>已上市<em class="vline"></em>1000-9999人</p>
                                            </div>
                                            <a href="/gongsi/f3e62b9ed1bfa70b1nNy09w~.html" ka="search_list_company_17_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/bar/20200629/7d893d759ab92844234643d8d033008dbe1bd4a3bd2a63f070bdbdada9aad826.png?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">gRPC</span>
                                                    <span class="tag-item">Gin</span>
                                                    <span class="tag-item">Gorm</span>
                                                    <span class="tag-item">架构师</span>
                                                    <span class="tag-item">Linux</span>
                                                    <span class="tag-item">后端开发</span>
                                                    <span class="tag-item">微服务架构</span>
                                                    <span class="tag-item">分布式技术</span>
                                        </div>
                                        <div class="info-desc">免费班车，餐补，补充医疗保险，带薪年假，五险一金，年终奖，定期体检，节日福利</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/762b1f113d53dbc71nJ92N-9E1ZU.html" data-securityid="lk_HTg_TO7Sk6-p1Ko9Fiij63V8uKB0mV12GK9BCg8uAoUYjJYOel63DNm_Mzu94QP9x7UsGgsqBp6m87Xkv1VRAgFufevR6m3PexPkbDABA4Ymk" data-jid="762b1f113d53dbc71nJ92N-9E1ZU" data-itemid="18" data-lid="9d7sWA96TcY.search.18" data-jobid="167320344" data-index="17" ka="search_list_18" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/762b1f113d53dbc71nJ92N-9E1ZU.html" title="Golang" target="_blank" ka="search_list_jname_18" data-securityid="lk_HTg_TO7Sk6-p1Ko9Fiij63V8uKB0mV12GK9BCg8uAoUYjJYOel63DNm_Mzu94QP9x7UsGgsqBp6m87Xkv1VRAgFufevR6m3PexPkbDABA4Ymk" data-jid="762b1f113d53dbc71nJ92N-9E1ZU" data-itemid="18" data-lid="9d7sWA96TcY.search.18" data-jobid="167320344" data-index="17">Golang</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·朝阳区·酒仙桥</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">20-40K</span>
                                                    <p>经验不限<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>彭先生<em class="vline"></em>架构师</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_18" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=lk_HTg_TO7Sk6-p1Ko9Fiij63V8uKB0mV12GK9BCg8uAoUYjJYOel63DNm_Mzu94QP9x7UsGgsqBp6m87Xkv1VRAgFufevR6m3PexPkbDABA4Ymk&jobId=762b1f113d53dbc71nJ92N-9E1ZU&lid=9d7sWA96TcY.search.18" redirect-url="/web/geek/chat?id=2e67aff3efa4035c1nd72tm-EFU~">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/5655835880426ec51nN62t-0.html" title="同程艺龙招聘" ka="search_list_company_18_custompage" target="_blank">同程艺龙</a></h3>
                                                <p><a href="/i100020/" class="false-link" target="_blank" ka="search_list_company_industry_18_custompage" title="互联网行业招聘信息">互联网</a><em class="vline"></em>已上市<em class="vline"></em>1000-9999人</p>
                                            </div>
                                            <a href="/gongsi/5655835880426ec51nN62t-0.html" ka="search_list_company_18_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/banner/12d614a4910d1b2aee2d6d8e408a402acfcd208495d565ef66e7dff9f98764da.jpg?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">容器技术</span>
                                                    <span class="tag-item">后端开发</span>
                                                    <span class="tag-item">服务端开发</span>
                                                    <span class="tag-item">Linux</span>
                                        </div>
                                        <div class="info-desc">定期体检，节日福利，五险一金，加班补助，通讯补贴，补充医疗保险，带薪年假，年终奖</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/acdad5e1d1c310b01nFy09y0FVRT.html" data-securityid="asTsAzXELrekZ-P1Ci0JMQrDjlbM3LaBQp2Xxg_BXqa8JWds-RPMjErpJ8pi_zjC3x3Uu3-2WpAR8cbolOMPUy4wXIBrZFXgtO5uPnoOuZwv6FI~" data-jid="acdad5e1d1c310b01nFy09y0FVRT" data-itemid="19" data-lid="9d7sWA96TcY.search.19" data-jobid="158819563" data-index="18" ka="search_list_19" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/acdad5e1d1c310b01nFy09y0FVRT.html" title="Golang" target="_blank" ka="search_list_jname_19" data-securityid="asTsAzXELrekZ-P1Ci0JMQrDjlbM3LaBQp2Xxg_BXqa8JWds-RPMjErpJ8pi_zjC3x3Uu3-2WpAR8cbolOMPUy4wXIBrZFXgtO5uPnoOuZwv6FI~" data-jid="acdad5e1d1c310b01nFy09y0FVRT" data-itemid="19" data-lid="9d7sWA96TcY.search.19" data-jobid="158819563" data-index="18">Golang</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·朝阳区·望京</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">25-45K·13薪</span>
                                                    <p>5-10年<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>邓先生<em class="vline"></em>架构师</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_19" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=asTsAzXELrekZ-P1Ci0JMQrDjlbM3LaBQp2Xxg_BXqa8JWds-RPMjErpJ8pi_zjC3x3Uu3-2WpAR8cbolOMPUy4wXIBrZFXgtO5uPnoOuZwv6FI~&jobId=acdad5e1d1c310b01nFy09y0FVRT&lid=9d7sWA96TcY.search.19" redirect-url="/web/geek/chat?id=d088d17dc0c4fc230nFy2dS0Ew~~">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/6e888338434f2e891nR52t-0.html" title="泡泡玛特POP MART招聘" ka="search_list_company_19_custompage" target="_blank">泡泡玛特POP MART</a></h3>
                                                <p><a href="/i100017/" class="false-link" target="_blank" ka="search_list_company_industry_19_custompage" title="新零售行业招聘信息">新零售</a><em class="vline"></em>已上市<em class="vline"></em>1000-9999人</p>
                                            </div>
                                            <a href="/gongsi/6e888338434f2e891nR52t-0.html" ka="search_list_company_19_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/bar/20200916/7ad8fb2f46da3bd1f3f367b5b7b5d564be1bd4a3bd2a63f070bdbdada9aad826.png?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">架构师</span>
                                                    <span class="tag-item">分布式技术</span>
                                                    <span class="tag-item">微服务架构</span>
                                                    <span class="tag-item">大数据</span>
                                                    <span class="tag-item">clickhouse</span>
                                                    <span class="tag-item">游戏后端</span>
                                        </div>
                                        <div class="info-desc">五险一金，定期体检，补充医疗保险，节日福利，餐补，交通补助，年终奖，带薪年假，零食下午茶，加班补助，员工旅游</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/2ba86de6139a35140Hx63t28FVU~.html" data-securityid="KmR1ZwmgjQDhD-c1pKkx9xpTG7HQRnYA9R_E9Ek3mfBTlbkC0chiWXbHcOtKczfO4IWybwi7f1k8cWOstH6lK2W6z0nt_4Ko82_V3Xq-h09QC74~" data-jid="2ba86de6139a35140Hx63t28FVU~" data-itemid="20" data-lid="9d7sWA96TcY.search.20" data-jobid="78050157" data-index="19" ka="search_list_20" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/2ba86de6139a35140Hx63t28FVU~.html" title="Golang" target="_blank" ka="search_list_jname_20" data-securityid="KmR1ZwmgjQDhD-c1pKkx9xpTG7HQRnYA9R_E9Ek3mfBTlbkC0chiWXbHcOtKczfO4IWybwi7f1k8cWOstH6lK2W6z0nt_4Ko82_V3Xq-h09QC74~" data-jid="2ba86de6139a35140Hx63t28FVU~" data-itemid="20" data-lid="9d7sWA96TcY.search.20" data-jobid="78050157" data-index="19">Golang</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·海淀区·上地</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">16-30K·15薪</span>
                                                    <p>经验不限<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>李先生<em class="vline"></em>云安全产品线副总经理</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_20" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=KmR1ZwmgjQDhD-c1pKkx9xpTG7HQRnYA9R_E9Ek3mfBTlbkC0chiWXbHcOtKczfO4IWybwi7f1k8cWOstH6lK2W6z0nt_4Ko82_V3Xq-h09QC74~&jobId=2ba86de6139a35140Hx63t28FVU~&lid=9d7sWA96TcY.search.20" redirect-url="/web/geek/chat?id=12584a5f771b01d603153dy-E1E~">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/84a0ab8293b2e3611Xd92d61.html" title="天融信招聘" ka="search_list_company_20_custompage" target="_blank">天融信</a></h3>
                                                <p><a href="/i100016/" class="false-link" target="_blank" ka="search_list_company_industry_20_custompage" title="信息安全行业招聘信息">信息安全</a><em class="vline"></em>已上市<em class="vline"></em>1000-9999人</p>
                                            </div>
                                            <a href="/gongsi/84a0ab8293b2e3611Xd92d61.html" ka="search_list_company_20_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/bar/20200916/487297720e21012062d5010b53c66384be1bd4a3bd2a63f070bdbdada9aad826.png?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">Go</span>
                                                    <span class="tag-item">容器技术</span>
                                                    <span class="tag-item">Linux</span>
                                                    <span class="tag-item">网络协议</span>
                                                    <span class="tag-item">gRPC</span>
                                                    <span class="tag-item">后端开发</span>
                                                    <span class="tag-item">服务端开发</span>
                                                    <span class="tag-item">管理系统</span>
                                        </div>
                                        <div class="info-desc">餐补，春节11天假，带薪年假，通讯补贴，定期体检，交通补助，节日福利，带薪病假5天，补充医疗保险，年终奖，股票期权，五险一金</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/4c55926927c6a9771nF80tm0ElRQ.html" data-securityid="sB_mZYaERW_Q2-x1k3PSCJTdpRV1QH0wxvQORMfXLHE5d6xI8Rp1mW4Gb_Fjh-lNYORacjLTyitn2h85a6PtOz6n6xWErrUiVnf66eDuzjgLQPnAOss~" data-jid="4c55926927c6a9771nF80tm0ElRQ" data-itemid="21" data-lid="9d7sWA96TcY.search.21" data-jobid="156949260" data-index="20" ka="search_list_21" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/4c55926927c6a9771nF80tm0ElRQ.html" title="Golang" target="_blank" ka="search_list_jname_21" data-securityid="sB_mZYaERW_Q2-x1k3PSCJTdpRV1QH0wxvQORMfXLHE5d6xI8Rp1mW4Gb_Fjh-lNYORacjLTyitn2h85a6PtOz6n6xWErrUiVnf66eDuzjgLQPnAOss~" data-jid="4c55926927c6a9771nF80tm0ElRQ" data-itemid="21" data-lid="9d7sWA96TcY.search.21" data-jobid="156949260" data-index="20">Golang</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·海淀区·双榆树</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">20-40K·15薪</span>
                                                    <p>经验不限<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>肖先生<em class="vline"></em>抖音电商服务端工程师</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_21" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=sB_mZYaERW_Q2-x1k3PSCJTdpRV1QH0wxvQORMfXLHE5d6xI8Rp1mW4Gb_Fjh-lNYORacjLTyitn2h85a6PtOz6n6xWErrUiVnf66eDuzjgLQPnAOss~&jobId=4c55926927c6a9771nF80tm0ElRQ&lid=9d7sWA96TcY.search.21" redirect-url="/web/geek/chat?id=4f93bffd70c025f80XR509y4GFI~">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/4cd971707633dd5d1nV82Nq6E1Y~.html" title="字节跳动招聘" ka="search_list_company_21_custompage" target="_blank">字节跳动</a></h3>
                                                <p><a href="/i100004/" class="false-link" target="_blank" ka="search_list_company_industry_21_custompage" title="广告营销行业招聘信息">广告营销</a><em class="vline"></em>不需要融资<em class="vline"></em>10000人以上</p>
                                            </div>
                                            <a href="/gongsi/4cd971707633dd5d1nV82Nq6E1Y~.html" ka="search_list_company_21_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/upload/com/logo/20210314/40431a563e843e1590d1f7d82652d53a0c6262faac2d042da229926610f4b0b6.png?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">Java</span>
                                                    <span class="tag-item">Shell</span>
                                                    <span class="tag-item">服务端开发</span>
                                        </div>
                                        <div class="info-desc">员工旅游，加班补助，餐补，住房补贴，通讯补贴，零食下午茶，节日福利，五险一金，带薪年假，定期体检，交通补助，年终奖，补充医疗保险，股票期权，包吃</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/cf39d6491d0fe5df1nJ93t27EFdX.html" data-securityid="IkfJipsL386YY-Y1AQ3Gh0-AsB2G8p1DxH5VzLJ52Caa71X171c1oACqU_D7bWNAo1UvSdpyQb2GwmtsLs9W3qArb_zgR8ZW2oNLV6VqWONuunX4UA~~" data-jid="cf39d6491d0fe5df1nJ93t27EFdX" data-itemid="22" data-lid="9d7sWA96TcY.search.22" data-jobid="167506057" data-index="21" ka="search_list_22" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/cf39d6491d0fe5df1nJ93t27EFdX.html" title="Golang" target="_blank" ka="search_list_jname_22" data-securityid="IkfJipsL386YY-Y1AQ3Gh0-AsB2G8p1DxH5VzLJ52Caa71X171c1oACqU_D7bWNAo1UvSdpyQb2GwmtsLs9W3qArb_zgR8ZW2oNLV6VqWONuunX4UA~~" data-jid="cf39d6491d0fe5df1nJ93t27EFdX" data-itemid="22" data-lid="9d7sWA96TcY.search.22" data-jobid="167506057" data-index="21">Golang</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·海淀区·中关村</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">25-26K</span>
                                                    <p>1年以内<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>张先生<em class="vline"></em>后端开发工程师</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_22" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=IkfJipsL386YY-Y1AQ3Gh0-AsB2G8p1DxH5VzLJ52Caa71X171c1oACqU_D7bWNAo1UvSdpyQb2GwmtsLs9W3qArb_zgR8ZW2oNLV6VqWONuunX4UA~~&jobId=cf39d6491d0fe5df1nJ93t27EFdX&lid=9d7sWA96TcY.search.22" redirect-url="/web/geek/chat?id=ca85c819b775337e0H1539y6FVo~">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/f409f37f83a6135b0nV_2d25EA~~.html" title="字节跳动招聘" ka="search_list_company_22_custompage" target="_blank">字节跳动</a></h3>
                                                <p><a href="/i100020/" class="false-link" target="_blank" ka="search_list_company_industry_22_custompage" title="互联网行业招聘信息">互联网</a><em class="vline"></em>D轮及以上<em class="vline"></em>10000人以上</p>
                                            </div>
                                            <a href="/gongsi/f409f37f83a6135b0nV_2d25EA~~.html" ka="search_list_company_22_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/upload/com/logo/20210323/aab1b65662224a865a9b6ba9aceaab02babfa420ecc134d2606b05aa8114130d.png?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">后端开发</span>
                                                    <span class="tag-item">架构师</span>
                                                    <span class="tag-item">Gin</span>
                                                    <span class="tag-item">分布式技术</span>
                                                    <span class="tag-item">微服务架构</span>
                                        </div>
                                        <div class="info-desc">年终奖，零食下午茶，包吃，加班补助，节日福利，住房补贴，补充医疗保险，带薪年假，交通补助，五险一金，定期体检，餐补</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/9ae1f9eb4a49e8dc1nF40t64EFdW.html" data-securityid="9vpam3JQ3LMMG-W1DsEE4YZr1giWnznPXzxgpp2J8Uiw0K06LmjKNxtyrrWbaG9uw8tNg0-G04zb3pSPvUi6qLvLdUJw1iV7o0Qt7NUf5AkZ" data-jid="9ae1f9eb4a49e8dc1nF40t64EFdW" data-itemid="23" data-lid="9d7sWA96TcY.search.23" data-jobid="152935056" data-index="22" ka="search_list_23" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/9ae1f9eb4a49e8dc1nF40t64EFdW.html" title="Golang" target="_blank" ka="search_list_jname_23" data-securityid="9vpam3JQ3LMMG-W1DsEE4YZr1giWnznPXzxgpp2J8Uiw0K06LmjKNxtyrrWbaG9uw8tNg0-G04zb3pSPvUi6qLvLdUJw1iV7o0Qt7NUf5AkZ" data-jid="9ae1f9eb4a49e8dc1nF40t64EFdW" data-itemid="23" data-lid="9d7sWA96TcY.search.23" data-jobid="152935056" data-index="22">Golang</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·海淀区·西北旺</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">25-40K·16薪</span>
                                                    <p>1-3年<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>彭先生<em class="vline"></em>后台开发工程师</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_23" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=9vpam3JQ3LMMG-W1DsEE4YZr1giWnznPXzxgpp2J8Uiw0K06LmjKNxtyrrWbaG9uw8tNg0-G04zb3pSPvUi6qLvLdUJw1iV7o0Qt7NUf5AkZ&jobId=9ae1f9eb4a49e8dc1nF40t64EFdW&lid=9d7sWA96TcY.search.23" redirect-url="/web/geek/chat?id=3d1390d35cbf26521XJ70tq1FFI~">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/2e64a887a110ea9f1nRz.html" title="腾讯招聘" ka="search_list_company_23_custompage" target="_blank">腾讯</a></h3>
                                                <p><a href="/i100020/" class="false-link" target="_blank" ka="search_list_company_industry_23_custompage" title="互联网行业招聘信息">互联网</a><em class="vline"></em>不需要融资<em class="vline"></em>10000人以上</p>
                                            </div>
                                            <a href="/gongsi/2e64a887a110ea9f1nRz.html" ka="search_list_company_23_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/bar/20200430/4204e9c9f200b00b77fb59d093acd281be1bd4a3bd2a63f070bdbdada9aad826.jpg?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">Go</span>
                                                    <span class="tag-item">PHP</span>
                                                    <span class="tag-item">SQL</span>
                                                    <span class="tag-item">分布式技术</span>
                                                    <span class="tag-item">后端开发</span>
                                        </div>
                                        <div class="info-desc">全勤奖，补充医疗保险，定期体检，零食下午茶，员工旅游，餐补，五险一金，免费班车，住房补贴，节日福利，带薪年假</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/659748d75068ac421nVz09S-FVFX.html" data-securityid="LyWkyKrstqyVS-71MfqHHUp8XLx8Wkgwq4MItWxQgCWzvZqf37SD0XDL_7_pdOviMzDBBlSmk690FIlY4eUsiVUTzV_XMPWAHCzo-RodRG58OkmHfg~~" data-jid="659748d75068ac421nVz09S-FVFX" data-itemid="24" data-lid="9d7sWA96TcY.search.24" data-jobid="119893537" data-index="23" ka="search_list_24" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/659748d75068ac421nVz09S-FVFX.html" title="Golang" target="_blank" ka="search_list_jname_24" data-securityid="LyWkyKrstqyVS-71MfqHHUp8XLx8Wkgwq4MItWxQgCWzvZqf37SD0XDL_7_pdOviMzDBBlSmk690FIlY4eUsiVUTzV_XMPWAHCzo-RodRG58OkmHfg~~" data-jid="659748d75068ac421nVz09S-FVFX" data-itemid="24" data-lid="9d7sWA96TcY.search.24" data-jobid="119893537" data-index="23">Golang</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·海淀区·牡丹园</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">30-60K·14薪</span>
                                                    <p>3-5年<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>包先生<em class="vline"></em>架构师</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_24" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=LyWkyKrstqyVS-71MfqHHUp8XLx8Wkgwq4MItWxQgCWzvZqf37SD0XDL_7_pdOviMzDBBlSmk690FIlY4eUsiVUTzV_XMPWAHCzo-RodRG58OkmHfg~~&jobId=659748d75068ac421nVz09S-FVFX&lid=9d7sWA96TcY.search.24" redirect-url="/web/geek/chat?id=00722967a33997d70nV429y0ElpZ">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/8f580ba1be2477691Xx70966.html" title="医渡云招聘" ka="search_list_company_24_custompage" target="_blank">医渡云</a></h3>
                                                <p><a href="/i100020/" class="false-link" target="_blank" ka="search_list_company_industry_24_custompage" title="互联网行业招聘信息">互联网</a><em class="vline"></em>已上市<em class="vline"></em>1000-9999人</p>
                                            </div>
                                            <a href="/gongsi/8f580ba1be2477691Xx70966.html" ka="search_list_company_24_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/bar/20201205/d2bc96af0f0807c709cebef3ae0d04d2be1bd4a3bd2a63f070bdbdada9aad826.png?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">分布式技术</span>
                                                    <span class="tag-item">网络编程</span>
                                                    <span class="tag-item">高级软件工程师</span>
                                                    <span class="tag-item">gRPC</span>
                                                    <span class="tag-item">golang</span>
                                                    <span class="tag-item">安全计算</span>
                                                    <span class="tag-item">区块链</span>
                                                    <span class="tag-item">数据结构</span>
                                        </div>
                                        <div class="info-desc">五险一金，股票期权，定期体检，节日福利，零食下午茶，年终奖，带薪年假，补充医疗保险，包吃</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/13ebaa14392912701nZ_09m7F1tX.html" data-securityid="rTnulOc4u4xQn-h176ZHTkztXpvYkFBG2af1B-j0qcL6TH1Hc9LplAzaC6vmQ6Z3bB7VUcudQOhwJe-8KHHqLczGPcoxShv1vdKNf3BYlPZBSoTFaQ~~" data-jid="13ebaa14392912701nZ_09m7F1tX" data-itemid="25" data-lid="9d7sWA96TcY.search.25" data-jobid="125846797" data-index="24" ka="search_list_25" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/13ebaa14392912701nZ_09m7F1tX.html" title="Golang" target="_blank" ka="search_list_jname_25" data-securityid="rTnulOc4u4xQn-h176ZHTkztXpvYkFBG2af1B-j0qcL6TH1Hc9LplAzaC6vmQ6Z3bB7VUcudQOhwJe-8KHHqLczGPcoxShv1vdKNf3BYlPZBSoTFaQ~~" data-jid="13ebaa14392912701nZ_09m7F1tX" data-itemid="25" data-lid="9d7sWA96TcY.search.25" data-jobid="125846797" data-index="24">Golang</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·大兴区·亦庄</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">25-50K</span>
                                                    <p>5-10年<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>杨先生<em class="vline"></em>高级软件开发工程师</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_25" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=rTnulOc4u4xQn-h176ZHTkztXpvYkFBG2af1B-j0qcL6TH1Hc9LplAzaC6vmQ6Z3bB7VUcudQOhwJe-8KHHqLczGPcoxShv1vdKNf3BYlPZBSoTFaQ~~&jobId=13ebaa14392912701nZ_09m7F1tX&lid=9d7sWA96TcY.search.25" redirect-url="/web/geek/chat?id=f04dd9fa1f7c14600nVz2ty9GFZY">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/33e052361693f8371nF-3d25.html" title="京东集团招聘" ka="search_list_company_25_custompage" target="_blank">京东集团</a></h3>
                                                <p><a href="/i100001/" class="false-link" target="_blank" ka="search_list_company_industry_25_custompage" title="电子商务行业招聘信息">电子商务</a><em class="vline"></em>已上市<em class="vline"></em>10000人以上</p>
                                            </div>
                                            <a href="/gongsi/33e052361693f8371nF-3d25.html" ka="search_list_company_25_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/bar/20191129/3cdf5ba2149e309b38868b62ae9c22cabe1bd4a3bd2a63f070bdbdada9aad826.jpg?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">分布式技术</span>
                                                    <span class="tag-item">微服务架构</span>
                                                    <span class="tag-item">后端开发</span>
                                                    <span class="tag-item">服务端开发</span>
                                                    <span class="tag-item">架构师</span>
                                                    <span class="tag-item">高级软件工程师</span>
                                        </div>
                                        <div class="info-desc">包吃，补充医疗保险，股票期权，年终奖，带薪年假，餐补，员工旅游，定期体检，交通补助，零食下午茶，节日福利，免费班车，五险一金</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/d6e11ce5b4593ed41nx42dW5EFNT.html" data-securityid="aGOC6tJ2yFpQC-I1iNK69puRwljuVCZn9VeLUfmhxazUbeXDwFi9mBRw6I1KsKcaNRxQoaC0UCNQjo-o5fT_JK8dVJbdccL3qpXWyPxfByzU50E~" data-jid="d6e11ce5b4593ed41nx42dW5EFNT" data-itemid="26" data-lid="9d7sWA96TcY.search.26" data-jobid="182284013" data-index="25" ka="search_list_26" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/d6e11ce5b4593ed41nx42dW5EFNT.html" title="Golang" target="_blank" ka="search_list_jname_26" data-securityid="aGOC6tJ2yFpQC-I1iNK69puRwljuVCZn9VeLUfmhxazUbeXDwFi9mBRw6I1KsKcaNRxQoaC0UCNQjo-o5fT_JK8dVJbdccL3qpXWyPxfByzU50E~" data-jid="d6e11ce5b4593ed41nx42dW5EFNT" data-itemid="26" data-lid="9d7sWA96TcY.search.26" data-jobid="182284013" data-index="25">Golang</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·海淀区·魏公村</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">25-50K·15薪</span>
                                                    <p>3-5年<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>杨先生<em class="vline"></em>RD</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_26" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=aGOC6tJ2yFpQC-I1iNK69puRwljuVCZn9VeLUfmhxazUbeXDwFi9mBRw6I1KsKcaNRxQoaC0UCNQjo-o5fT_JK8dVJbdccL3qpXWyPxfByzU50E~&jobId=d6e11ce5b4593ed41nx42dW5EFNT&lid=9d7sWA96TcY.search.26" redirect-url="/web/geek/chat?id=4613f089597e5a3e33Fy3du9EFE~">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/a67b361452e384e71XV82N4~.html" title="今日头条招聘" ka="search_list_company_26_custompage" target="_blank">今日头条</a></h3>
                                                <p><a href="/i100019/" class="false-link" target="_blank" ka="search_list_company_industry_26_custompage" title="移动互联网行业招聘信息">移动互联网</a><em class="vline"></em>不需要融资<em class="vline"></em>10000人以上</p>
                                            </div>
                                            <a href="/gongsi/a67b361452e384e71XV82N4~.html" ka="search_list_company_26_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/upload/com/logo/20210525/77d60eae41e48b90df64951371a7a07a19f97e2c258c6cead07beaf11928d91b.png?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">Java</span>
                                                    <span class="tag-item">PHP</span>
                                                    <span class="tag-item">Python</span>
                                                    <span class="tag-item">数据结构</span>
                                                    <span class="tag-item">后端开发</span>
                                                    <span class="tag-item">Linux</span>
                                        </div>
                                        <div class="info-desc">免费班车，通讯补贴，年终奖，定期体检，补充医疗保险，节日福利，住房补贴，零食下午茶，五险一金，六险一金，交通补助，员工旅游，试用期同薪，健身房，加班补助，带薪年假，股票期权，六险一金</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/3571edfada1473b71nJ80t26E1tZ.html" data-securityid="1ZtWeu0ZYtnyD-g1VUFQXPcAqTQXWBYv1xoOgGUjaL-yEakjRceMc5y2ykaCBAMgRwhsJ0GZloDsqY4qfdIlpkqMd_KPB_sQ_haYJqqN-34~" data-jid="3571edfada1473b71nJ80t26E1tZ" data-itemid="27" data-lid="9d7sWA96TcY.search.27" data-jobid="166907399" data-index="26" ka="search_list_27" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/3571edfada1473b71nJ80t26E1tZ.html" title="Golang" target="_blank" ka="search_list_jname_27" data-securityid="1ZtWeu0ZYtnyD-g1VUFQXPcAqTQXWBYv1xoOgGUjaL-yEakjRceMc5y2ykaCBAMgRwhsJ0GZloDsqY4qfdIlpkqMd_KPB_sQ_haYJqqN-34~" data-jid="3571edfada1473b71nJ80t26E1tZ" data-itemid="27" data-lid="9d7sWA96TcY.search.27" data-jobid="166907399" data-index="26">Golang</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·海淀区·上地</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">20-40K</span>
                                                    <p>经验不限<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>吴女士<em class="vline"></em>HR</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_27" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=1ZtWeu0ZYtnyD-g1VUFQXPcAqTQXWBYv1xoOgGUjaL-yEakjRceMc5y2ykaCBAMgRwhsJ0GZloDsqY4qfdIlpkqMd_KPB_sQ_haYJqqN-34~&jobId=3571edfada1473b71nJ80t26E1tZ&lid=9d7sWA96TcY.search.27" redirect-url="/web/geek/chat?id=b5024faf27010cdf0Xxy0tu0EFo~">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/ab9fdc6f043679990HY~.html" title="百度招聘" ka="search_list_company_27_custompage" target="_blank">百度</a></h3>
                                                <p><a href="/i100020/" class="false-link" target="_blank" ka="search_list_company_industry_27_custompage" title="互联网行业招聘信息">互联网</a><em class="vline"></em>已上市<em class="vline"></em>10000人以上</p>
                                            </div>
                                            <a href="/gongsi/ab9fdc6f043679990HY~.html" ka="search_list_company_27_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/logo/0e0d441a2f93a236536f20e8277bf4dbbe1bd4a3bd2a63f070bdbdada9aad826.png?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">PHP</span>
                                        </div>
                                        <div class="info-desc">餐补，团队氛围好，零食下午茶，带薪年假，老板Nice，加班补助，住房补贴，娱乐健身，五险一金，节日福利，员工旅游，定期体检，交通补助，年终奖，补充医疗保险，通讯补贴，免费班车</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/b95e9d9e3030404c1nJ43t-0E1JY.html" data-securityid="snpNqiKwhryuP-q1hL0xoVjRzgpYbM4WQfGZTpERvjxnu-hXRq23a6OnNzRPp9R6oaJJZ4rPQBE1OVgnS2bnWhmGqh5-KQNSEW78EkzDjICOHULk8PQ~" data-jid="b95e9d9e3030404c1nJ43t-0E1JY" data-itemid="28" data-lid="9d7sWA96TcY.search.28" data-jobid="162529308" data-index="27" ka="search_list_28" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/b95e9d9e3030404c1nJ43t-0E1JY.html" title="Golang" target="_blank" ka="search_list_jname_28" data-securityid="snpNqiKwhryuP-q1hL0xoVjRzgpYbM4WQfGZTpERvjxnu-hXRq23a6OnNzRPp9R6oaJJZ4rPQBE1OVgnS2bnWhmGqh5-KQNSEW78EkzDjICOHULk8PQ~" data-jid="b95e9d9e3030404c1nJ43t-0E1JY" data-itemid="28" data-lid="9d7sWA96TcY.search.28" data-jobid="162529308" data-index="27">Golang</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·海淀区·中关村</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">30-60K·15薪</span>
                                                    <p>经验不限<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>朱先生<em class="vline"></em>抖音UG服务端研发</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_28" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=snpNqiKwhryuP-q1hL0xoVjRzgpYbM4WQfGZTpERvjxnu-hXRq23a6OnNzRPp9R6oaJJZ4rPQBE1OVgnS2bnWhmGqh5-KQNSEW78EkzDjICOHULk8PQ~&jobId=b95e9d9e3030404c1nJ43t-0E1JY&lid=9d7sWA96TcY.search.28" redirect-url="/web/geek/chat?id=dce5b483e142350e0nd83tW9FlZR">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/1311d10fe26c4ee51Hd73ti0FA~~.html" title="字节跳动招聘" ka="search_list_company_28_custompage" target="_blank">字节跳动</a></h3>
                                                <p><a href="/i100019/" class="false-link" target="_blank" ka="search_list_company_industry_28_custompage" title="移动互联网行业招聘信息">移动互联网</a><em class="vline"></em>D轮及以上<em class="vline"></em>10000人以上</p>
                                            </div>
                                            <a href="/gongsi/1311d10fe26c4ee51Hd73ti0FA~~.html" ka="search_list_company_28_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/chatphoto/20201123/6a4f3bc17c6a4f8fea5a9ba01cb22bcd82e5997348729a6cfe816ad90c892e55_s.jpg?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">后端开发</span>
                                        </div>
                                        <div class="info-desc">餐补，交通补助，节日福利</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/0b7264384a3a1b831nJ53d-6FlZZ.html" data-securityid="nLWa9ONBprBOe-8122rHCgt1rGVMTNUdUrQbq5somBegNIuXP6LSzkbSZC0kpBImZA-wBrlIGodVD61w7ZKe1DkjSKqd34ctSvtSIV9K9ZxKB6U~" data-jid="0b7264384a3a1b831nJ53d-6FlZZ" data-itemid="29" data-lid="9d7sWA96TcY.search.29" data-jobid="163627649" data-index="28" ka="search_list_29" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/0b7264384a3a1b831nJ53d-6FlZZ.html" title="Golang" target="_blank" ka="search_list_jname_29" data-securityid="nLWa9ONBprBOe-8122rHCgt1rGVMTNUdUrQbq5somBegNIuXP6LSzkbSZC0kpBImZA-wBrlIGodVD61w7ZKe1DkjSKqd34ctSvtSIV9K9ZxKB6U~" data-jid="0b7264384a3a1b831nJ53d-6FlZZ" data-itemid="29" data-lid="9d7sWA96TcY.search.29" data-jobid="163627649" data-index="28">Golang</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·朝阳区·望京</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">15-25K</span>
                                                    <p>3-5年<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>张女士<em class="vline"></em>hrbp</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_29" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=nLWa9ONBprBOe-8122rHCgt1rGVMTNUdUrQbq5somBegNIuXP6LSzkbSZC0kpBImZA-wBrlIGodVD61w7ZKe1DkjSKqd34ctSvtSIV9K9ZxKB6U~&jobId=0b7264384a3a1b831nJ53d-6FlZZ&lid=9d7sWA96TcY.search.29" redirect-url="/web/geek/chat?id=792b7a80f5ec7de51HN72dW8GVQ~">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/f0682137f4215b9c1nd_3t0~.html" title="容联云通讯招聘" ka="search_list_company_29_custompage" target="_blank">容联云通讯</a></h3>
                                                <p><a href="/i100020/" class="false-link" target="_blank" ka="search_list_company_industry_29_custompage" title="互联网行业招聘信息">互联网</a><em class="vline"></em>已上市<em class="vline"></em>1000-9999人</p>
                                            </div>
                                            <a href="/gongsi/f0682137f4215b9c1nd_3t0~.html" ka="search_list_company_29_custompage_logo" target="_blank"><img class="company-logo" src="https://img2.bosszhipin.com/mcs/chatphoto/20151217/ea05d4b7e159211aeac51149edf5cf7f36ca6403f0c8c79246e0849b17594ff7.jpg?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">Python</span>
                                                    <span class="tag-item">分布式技术</span>
                                                    <span class="tag-item">容器技术</span>
                                                    <span class="tag-item">消息队列</span>
                                                    <span class="tag-item">后端开发</span>
                                                    <span class="tag-item">服务端开发</span>
                                        </div>
                                        <div class="info-desc">定期体检，五险一金，带薪年假，员工旅游，补充医疗保险，零食下午茶，节日福利，餐补，交通补助，加班补助，通讯补贴</div>
                                    </div>
                                </div>
                            </li>
                            <li>
                                <div class="job-primary">
                                    
                                    
                                    
                                    <div class="info-primary">
                                        <div class="primary-wrapper">
                                            <div class="primary-box" href="/job_detail/f1c0b4c8402bbbc51nZy3Nu4FVpW.html" data-securityid="h7jPO52R2LdGj-L1rQC6-oq4Z7BhhTIssNdqV9Bn2Rjjlz4loeCWdnpShy3P9YZhZoc20Y_RKBT9pS6SbR1ywWSZoLlYyZHDrdamC-5RfSpDIfM~" data-jid="f1c0b4c8402bbbc51nZy3Nu4FVpW" data-itemid="30" data-lid="9d7sWA96TcY.search.30" data-jobid="128765586" data-index="29" ka="search_list_30" target="_blank">
                                                <div class="job-title">
                                                    <span class="job-name"><a href="/job_detail/f1c0b4c8402bbbc51nZy3Nu4FVpW.html" title="Golang" target="_blank" ka="search_list_jname_30" data-securityid="h7jPO52R2LdGj-L1rQC6-oq4Z7BhhTIssNdqV9Bn2Rjjlz4loeCWdnpShy3P9YZhZoc20Y_RKBT9pS6SbR1ywWSZoLlYyZHDrdamC-5RfSpDIfM~" data-jid="f1c0b4c8402bbbc51nZy3Nu4FVpW" data-itemid="30" data-lid="9d7sWA96TcY.search.30" data-jobid="128765586" data-index="29">Golang</a></span>
                                                    <span class="job-area-wrapper">
                                                        <span class="job-area">北京·海淀区·人民大学</span>
                                                    </span>
                                                    <span class="job-pub-time"></span>
                                                </div>
                                                
                                                <div class="job-limit clearfix">
                                                    <span class="red">30-45K</span>
                                                    <p>1-3年<em class="vline"></em>本科</p>
                                                    <div class="info-publis">
                                                        <h3 class="name"><img class="icon-chat" src="https://z.zhipin.com/web/geek/resource/icon-chat-v2.png"/>李女士<em class="vline"></em>招聘经理</h3>
                                                    </div>
                                                    <button ka="cpc_job_list_chat_30" class="btn btn-startchat" href="javascript:;" data-url="/wapi/zpgeek/friend/add.json?securityId=h7jPO52R2LdGj-L1rQC6-oq4Z7BhhTIssNdqV9Bn2Rjjlz4loeCWdnpShy3P9YZhZoc20Y_RKBT9pS6SbR1ywWSZoLlYyZHDrdamC-5RfSpDIfM~&jobId=f1c0b4c8402bbbc51nZy3Nu4FVpW&lid=9d7sWA96TcY.search.30" redirect-url="/web/geek/chat?id=96d49552c873ca941HF53tu6EA~~">
                                                        <img class="icon-chat icon-chat-hover" src="https://z.zhipin.com/web/geek/resource/icon-chat-hover-v2.png" alt="">
                                                        <span>立即沟通</span>
                                                    </button>
                                                </div>
                                                <div class="info-detail"></div>
                                            </div>
                                        </div>
                                        <div class="info-company">
                                            <div class="company-text">
                                                <h3 class="name"><a href="/gongsi/e77fb14bf749756b1nZ-2tm5.html" title="微步在线招聘" ka="search_list_company_30_custompage" target="_blank">微步在线</a></h3>
                                                <p><a href="/i100016/" class="false-link" target="_blank" ka="search_list_company_industry_30_custompage" title="信息安全行业招聘信息">信息安全</a><em class="vline"></em>D轮及以上<em class="vline"></em>500-999人</p>
                                            </div>
                                            <a href="/gongsi/e77fb14bf749756b1nZ-2tm5.html" ka="search_list_company_30_custompage_logo" target="_blank"><img class="company-logo" src="https://img.bosszhipin.com/beijin/mcs/chatphoto/20200317/f7feec1da30b7cdd2a52cab334cce8755d616a4dad422782a68f228e273ba898_s.jpg?x-oss-process=image/resize,w_100,limit_0" alt=""></a>
                                        </div>
                                    </div>
                                    <div class="info-append clearfix">
                                        <div class="tags">
                                                    <span class="tag-item">分布式技术</span>
                                                    <span class="tag-item">网络协议</span>
                                                    <span class="tag-item">网络编程</span>
                                                    <span class="tag-item">多线程</span>
                                                    <span class="tag-item">后端开发</span>
                                                    <span class="tag-item">Linux</span>
                                                    <span class="tag-item">Unix</span>
                                                    <span class="tag-item">架构师</span>
                                        </div>
                                        <div class="info-desc">交通补助，股票期权，节日福利，司龄假，餐补，零食下午茶，员工旅游，包吃，补充医疗保险，五险一金，租房补贴，定期体检，年终奖，带薪年假</div>
                                    </div>
                                </div>
                            </li>
                    </ul>
        <div class="page">
            <a href="/c101010100-p100199/?query=golang&page=0" ka="page-prev" class="prev"></a>
        
        
        
        
        <a href="javascript:;" class="cur" ka="page-cur">1</a>
        <a href="/c101010100-p100199/?query=golang&page=2" ka="page-2">2</a>
        <a href="/c101010100-p100199/?query=golang&page=3" ka="page-3">3</a>
        <span>...</span>
            <a href="/c101010100-p100199/?query=golang&page=2" ka="page-next" class="next"></a>
        </div>
            </div>
                <form class="satisfaction-feedback">
                    <b class="title">对搜索结果是否满意？</b>
                    <div class="satisfaction level-2" data-level="1">不满意</div>
                    <div class="satisfaction level-3" data-level="2">一般</div>
                    <div class="satisfaction level-4" data-level="3">满意</div>
                    <textarea class="ipt" name="feedback" placeholder="请填写更多反馈建议..."></textarea>
                    <button class="btn disabled" disabled="disabled">提交</button>
                    <input type="hidden" class="form-data" name="level" value=""/>
                    <input type="hidden" class="form-data" name="query" value="golang"/>
                    <input type="hidden" class="form-data" name="c" value="101010100"/>
                    <input type="hidden" class="form-data" name="i" value=""/>
                    <input type="hidden" class="form-data" name="p" value="100199"/>
                    <input type="hidden" class="form-data" name="e" value=""/>
                    <input type="hidden" class="form-data" name="d" value=""/>
                    <input type="hidden" class="form-data" name="s" value=""/>
                    <input type="hidden" class="form-data" name="y" value=""/>
                    <input type="hidden" class="form-data" name="t" value=""/>
                    <input type="hidden" class="form-data" name="a" value=""/>
                    <input type="hidden" class="form-data" name="b" value=""/>
                    <input type="hidden" class="form-data" name="lid" value="9d7sWA96TcY"/>
                </form>
            <div class="pos-bread">
                    <a href="/beijing/" title="北京BOSS直聘招聘" ka="job-breadcrumb-bottom-0">北京BOSS直聘招聘</a>&gt;
                    <a href="/c101010100/" title="北京招聘" ka="job-breadcrumb-bottom-1">北京招聘</a>&gt;
                                        <h1 class="link-redirect"><a href="/c101010100-p100199/" title="北京后端开发招聘" ka="job-breadcrumb-bottom-2">北京后端开发招聘</a></h1>
            </div>
        </div>
    </div>
<div class="links-container ">
    <div class="links-box">
        <div class="title-box">
            <span class="title-item cur">热门职位</span>
            <span class="title-item">热门城市</span>
            <span class="title-item">热门企业</span>
            <span class="title-item">附近城市</span>
        </div>
        <div class="links-wrapper">
            <div class="links-content cur">
                <a href="/c101010100-p100101/" target="_blank" ka="hot-job-0">北京Java招聘</a>
                <a href="/c101010100-p100102/" target="_blank" ka="hot-job-1">北京C++招聘</a>
                <a href="/c101010100-p100103/" target="_blank" ka="hot-job-2">北京PHP招聘</a>
                <a href="/c101010100-p100105/" target="_blank" ka="hot-job-3">北京C招聘</a>
                <a href="/c101010100-p100106/" target="_blank" ka="hot-job-4">北京C#招聘</a>
                <a href="/c101010100-p100107/" target="_blank" ka="hot-job-5">北京.NET招聘</a>
                <a href="/c101010100-p100108/" target="_blank" ka="hot-job-6">北京Hadoop招聘</a>
                <a href="/c101010100-p100109/" target="_blank" ka="hot-job-7">北京Python招聘</a>
                <a href="/c101010100-p100110/" target="_blank" ka="hot-job-8">北京Delphi招聘</a>
                <a href="/c101010100-p100111/" target="_blank" ka="hot-job-9">北京VB招聘</a>
                <a href="/c101010100-p100112/" target="_blank" ka="hot-job-10">北京Perl招聘</a>
                <a href="/c101010100-p100113/" target="_blank" ka="hot-job-11">北京Ruby招聘</a>
                <a href="/c101010100-p100114/" target="_blank" ka="hot-job-12">北京Node.js招聘</a>
                <a href="/c101010100-p100116/" target="_blank" ka="hot-job-13">北京Golang招聘</a>
                <a href="/c101010100-p100119/" target="_blank" ka="hot-job-14">北京Erlang招聘</a>
                <a href="/c101010100-p100121/" target="_blank" ka="hot-job-15">北京语音/视频/图形开发招聘</a>
                <a href="/c101010100-p100122/" target="_blank" ka="hot-job-16">北京数据采集招聘</a>
                <a href="/c101010100-p100199/" target="_blank" ka="hot-job-17">北京后端开发招聘</a>
                <a href="/c101010100-p100123/" target="_blank" ka="hot-job-18">北京全栈工程师招聘</a>
                <a href="/c101010100-p100124/" target="_blank" ka="hot-job-19">北京GIS工程师招聘</a>
            </div>
            <div class="links-content">
                <a href="/beijing/" target="_blank" ka="hot-city-0">北京招聘</a>
                <a href="/shanghai/" target="_blank" ka="hot-city-1">上海招聘</a>
                <a href="/guangzhou/" target="_blank" ka="hot-city-2">广州招聘</a>
                <a href="/shenzhen/" target="_blank" ka="hot-city-3">深圳招聘</a>
                <a href="/hangzhou/" target="_blank" ka="hot-city-4">杭州招聘</a>
                <a href="/tianjin/" target="_blank" ka="hot-city-5">天津招聘</a>
                <a href="/xian/" target="_blank" ka="hot-city-6">西安招聘</a>
                <a href="/suzhou/" target="_blank" ka="hot-city-7">苏州招聘</a>
                <a href="/wuhan/" target="_blank" ka="hot-city-8">武汉招聘</a>
                <a href="/xiamen/" target="_blank" ka="hot-city-9">厦门招聘</a>
                <a href="/changsha/" target="_blank" ka="hot-city-10">长沙招聘</a>
                <a href="/chengdu/" target="_blank" ka="hot-city-11">成都招聘</a>
                <a href="/zhengzhou/" target="_blank" ka="hot-city-12">郑州招聘</a>
                <a href="/chongqing/" target="_blank" ka="hot-city-13">重庆招聘</a>
                <a href="/foshan/" target="_blank" ka="hot-city-14">佛山招聘</a>
                <a href="/hefei/" target="_blank" ka="hot-city-15">合肥招聘</a>
                <a href="/jinan/" target="_blank" ka="hot-city-16">济南招聘</a>
                <a href="/qingdao/" target="_blank" ka="hot-city-17">青岛招聘</a>
                <a href="/nanjing/" target="_blank" ka="hot-city-18">南京招聘</a>
                <a href="/dongguan/" target="_blank" ka="hot-city-19">东莞招聘</a>
                <a href="/kunming/" target="_blank" ka="hot-city-20">昆明招聘</a>
                <a href="/nanchang/" target="_blank" ka="hot-city-21">南昌招聘</a>
                <a href="/shijiazhuang/" target="_blank" ka="hot-city-22">石家庄招聘</a>
                <a href="/ningbo/" target="_blank" ka="hot-city-23">宁波招聘</a>
                <a href="/fuzhou/" target="_blank" ka="hot-city-24">福州招聘</a>
            </div>
            <div class="links-content">
                <a href="/gongsi/ab9fdc6f043679990HY~.html" target="_blank" ka="hot-brand-0">百度</a>
                <a href="/gongsi/480261c022ea03d81nV53tQ~.html" target="_blank" ka="hot-brand-1">快手</a>
                <a href="/gongsi/5d627415a46b4a750nJ9.html" target="_blank" ka="hot-brand-2">阿里巴巴集团</a>
                <a href="/gongsi/fa2f92669c66eee31Hc~.html" target="_blank" ka="hot-brand-3">BOSS直聘</a>
                <a href="/gongsi/2e64a887a110ea9f1nRz.html" target="_blank" ka="hot-brand-4">腾讯</a>
                <a href="/gongsi/efa878c051261c7433Z6.html" target="_blank" ka="hot-brand-5">好未来</a>
                <a href="/gongsi/01da85cd2b2d314a1HN40w~~.html" target="_blank" ka="hot-brand-6">高德地图</a>
                <a href="/gongsi/02cd05cce753437e33V50w~~.html" target="_blank" ka="hot-brand-7">华为</a>
                <a href="/gongsi/980f48937a13792b1nd63d0~.html" target="_blank" ka="hot-brand-8">滴滴出行</a>
                <a href="/gongsi/92f44a5f422a7d6f1nVy2Nk~.html" target="_blank" ka="hot-brand-9">搜狗</a>
                <a href="/gongsi/aad0931d89e8d9371nR-3Q~~.html" target="_blank" ka="hot-brand-10">好大夫在线</a>
                <a href="/gongsi/38bd5c757efa4ab6331z.html" target="_blank" ka="hot-brand-11">网易</a>
                <a href="/gongsi/d063e597d01c108b33Z50g~~.html" target="_blank" ka="hot-brand-12">亚信科技</a>
                <a href="/gongsi/2d208a8834e4a58203d43Q~~.html" target="_blank" ka="hot-brand-13">中软国际</a>
                <a href="/gongsi/12e92393bff967950HNy3w~~.html" target="_blank" ka="hot-brand-14">饿了么</a>
                <a href="/gongsi/0bc562732fcf91fe3nJy.html" target="_blank" ka="hot-brand-15">货拉拉科技</a>
                <a href="/gongsi/fd52e169fcf6f5591nZy2tg~.html" target="_blank" ka="hot-brand-16">汽车之家</a>
                <a href="/gongsi/139ba6a6fa587b411nd629w~.html" target="_blank" ka="hot-brand-17">完美世界</a>
                <a href="/gongsi/133fca828487ac3d1nV52A~~.html" target="_blank" ka="hot-brand-18">一起教育科技</a>
                <a href="/gongsi/7a2b06c8787b8fb30nBz2A~~.html" target="_blank" ka="hot-brand-19">一点资讯</a>
                <a href="/gongsi/ebad8a37af8239c11Xx92Q~~.html" target="_blank" ka="hot-brand-20">VIPKID</a>
                <a href="/gongsi/0172fe3e5bdbbb4c3nd93w~~.html" target="_blank" ka="hot-brand-21">英雄互娱</a>
                <a href="/gongsi/8531731ebc286f821nV73t4~.html" target="_blank" ka="hot-brand-22">京北方</a>
                <a href="/gongsi/0ac3daa93b4ee7011XN4.html" target="_blank" ka="hot-brand-23">奇虎360</a>
                <a href="/gongsi/546490316d69463a03J72A~~.html" target="_blank" ka="hot-brand-24">易车</a>
                <a href="/gongsi/e4b2c733e23c06761nZ53N0~.html" target="_blank" ka="hot-brand-25">中科软</a>
                <a href="/gongsi/6cd250243c0c51681nRz3dw~.html" target="_blank" ka="hot-brand-26">Keep</a>
                <a href="/gongsi/15be1afc3ceda8113nNy2w~~.html" target="_blank" ka="hot-brand-27">美术宝</a>
                <a href="/gongsi/075e0d17cd137e971nV43NU~.html" target="_blank" ka="hot-brand-28">旷视MEGVII</a>
                <a href="/gongsi/4d91ea8ded8f515c0nZ43A~~.html" target="_blank" ka="hot-brand-29">永航科技</a>
                <a href="/gongsi/f470bd8b47fbc1821nd709U~.html" target="_blank" ka="hot-brand-30">金山云</a>
                <a href="/gongsi/a4643c0fa9f97efc3nx_2A~~.html" target="_blank" ka="hot-brand-31">全时天地在线</a>
                <a href="/gongsi/0979070fc97e501a0nx7.html" target="_blank" ka="hot-brand-32">美团网</a>
                <a href="/gongsi/634abae943825f8c1H193g~~.html" target="_blank" ka="hot-brand-33">公瑾科技</a>
                <a href="/gongsi/0ed7168c0705429b1nZ-3g~~.html" target="_blank" ka="hot-brand-34">掌阅科技</a>
                <a href="/gongsi/ae1ff0467cbd29ea1nZ-3ts~.html" target="_blank" ka="hot-brand-35">bilibili</a>
                <a href="/gongsi/9cc010209b81ffb71nV-290~.html" target="_blank" ka="hot-brand-36">四维图新</a>
                <a href="/gongsi/8d038826246b5fb31nV83d4~.html" target="_blank" ka="hot-brand-37">泛微网络</a>
                <a href="/gongsi/918159f26789c3891nV53dQ~.html" target="_blank" ka="hot-brand-38">小红书</a>
                <a href="/gongsi/5506314c3728d75b1HJ60w~~.html" target="_blank" ka="hot-brand-39">致远互联</a>
            </div>
            <div class="links-content">
                <a href="/c101090100-p100199/" target="_blank" ka="nearby-city-0">石家庄后端开发招聘信息</a>
                <a href="/c101090200-p100199/" target="_blank" ka="nearby-city-1">保定后端开发招聘信息</a>
                <a href="/c101090300-p100199/" target="_blank" ka="nearby-city-2">张家口后端开发招聘信息</a>
                <a href="/c101090400-p100199/" target="_blank" ka="nearby-city-3">承德后端开发招聘信息</a>
                <a href="/c101090500-p100199/" target="_blank" ka="nearby-city-4">唐山后端开发招聘信息</a>
                <a href="/c101090600-p100199/" target="_blank" ka="nearby-city-5">廊坊后端开发招聘信息</a>
                <a href="/c101090700-p100199/" target="_blank" ka="nearby-city-6">沧州后端开发招聘信息</a>
                <a href="/c101090800-p100199/" target="_blank" ka="nearby-city-7">衡水后端开发招聘信息</a>
                <a href="/c101090900-p100199/" target="_blank" ka="nearby-city-8">邢台后端开发招聘信息</a>
                <a href="/c101091000-p100199/" target="_blank" ka="nearby-city-9">邯郸后端开发招聘信息</a>
                <a href="/c101091100-p100199/" target="_blank" ka="nearby-city-10">秦皇岛后端开发招聘信息</a>
            </div>
        </div>
        <div class="expand-btn"><a href="javascript:;" class="more-view"><span>展开</span><i class="fz fz-slidedown"></i></a></div>
    </div>
    <div class="links-box-new"><h3>北京后端开发招聘频道介绍</h3><p>BOSS直聘北京后端开发招聘频道，展示北京后端开发招聘信息，百万Boss在线直聘，直接开聊，在线面试，北京找后端开发工作就上BOSS直聘网站或APP！</p ></div>
</div>
    <!--底部footer-->
<div id="footer" class="">
    <div class="inner home-inner">
        <div class="footer-social">
            <img src="https://static.zhipin.com/v2/web/geek/images/footer-logo.png" alt="">
            <p>企业服务热线和举报投诉/未成年人举报渠道 <span>400 065 5799</span></p>
            <p>工作日 <span>8:00 - 22:00</span></p>
            <p>休息日 <span>9:30 - 18:30</span></p>
            <p class="footer-icon">
                <a href="https://www.weibo.com/bosszhipin" rel="nofollow" ka="link-weibo" target="_blank" class="icon-weibo"><span>官方微博</span></a>
                <a href="javascript:;" class="icon-weixin" ka="link-weixin"><span>微信公众号</span><img src="https://static.zhipin.com/v2/web/geek/images/we_chat_qr.jpg" class="qrcode-weixin" /></a>
                <a href="https://app.zhipin.com" class="icon-app" ka="link-downloadapp"><span>下载</span></a>
            </p>
        </div>
        <div class="footer-about clear-fix">
            <dl>
                <dt>企业服务</dt>
                <dd>
                    <a href="https://www.zhipin.com/job_detail/" ka="link-search">职位搜索</a>
                    <a href="https://news.zhipin.com/" ka="link-news">新闻资讯</a>
                    <a href="https://app.zhipin.com/" ka="link-app">BOSS直聘APP</a>
                    <a href="https://ir.zhipin.com/" target="_blank" ka="link-ir">投资者关系</a>
                </dd>
            </dl>
            <dl>
                <dt>使用与帮助</dt>
                <dd>
                    <a href="https://about.zhipin.com/agreement" rel="nofollow" ka="link-privacy" target="_blank">协议与规则</a>
                    <a href="https://about.zhipin.com/agreement?id=personalinfopro" rel="nofollow" ka="link-privacy" target="_blank">个人信息保护政策</a>
                    <a href="https://www.zhipin.com/activity/cc/anticheatguide.html" rel="nofollow" ka="link-anticheatguide" target="_blank">防骗指南</a>
                    <a href="https://www.zhipin.com/activity/cc/usehelp.html" rel="nofollow" ka="link-usehelp" target="_blank">使用帮助</a>
                </dd>
            </dl>
            <dl>
                <dt>联系BOSS直聘</dt>
                <dd>
                    <p>北京华品博睿网络技术有限公司</p>
                    <p>公司地址 北京市朝阳区太阳宫中路16号院1号楼18层1801内09</p>
                    <p>违法和不良信息举报邮箱/未成年人举报渠道 <a class="report-mail" href="mailto:jubao@kanzhun.com" rel="nofollow">jubao@kanzhun.com</a></p>
                </dd>
            </dl>
        </div>
        <div class="copyright">
            <p>
                <span>Copyright © 2021 BOSS直聘</span><span><a href="https://beian.miit.gov.cn/" rel="nofollow" ka="link-icp" target="_blank">京ICP备14013441号-5</a></span><span>京ICP证150923号</span><span>京网文[2020]0399-066 号</span>
                <span>
                     <a href="https://www.zhipin.com/activity/cc/businesslicense.html" rel="nofollow" ka="link-businesslicense" target="_blank">
                         <img src="https://static.zhipin.com/v2/web/geek/images/icon-badge-1.png" alt=""/>
                         电子营业执照
                     </a>
                </span>
                <span><a href="http://www.beian.gov.cn/portal/registerSystemInfo?recordcode=11010502032801" rel="nofollow" ka="link-beian" target="_blank"><img src="https://static.zhipin.com/v2/web/geek/images/icon-beian.png" alt=""/> 京公网安备11010502032801</a></span>
                <span><a href="http://sdwj.zhipin.com/web/index.html" rel="nofollow" ka="link-sdwj" target="_blank">朝阳网警</a></span>
                <span><a href="https://www.zhipin.com/web/common/protocol/hr-service.html" rel="nofollow" ka="link-hr-service" target="_blank">人力资源服务许可证</a></span>
                <span><a href="https://www.12377.cn" rel="nofollow" ka="link-12377" target="_blank">网上有害信息举报专区</a></span>
                <span class="renshe-phone">朝阳区人社局监督电话：(010)57596212，(010)65090445</span>
            </p>
        </div>
    </div>
</div>
    <!-- 通用侧边栏 -->
    <div id="siderbar">
        <div class="sider-index">
            <ul class="siderbar-top">
                <li data-value="interest">
                    <a href="javascript:;" ka="side_interest"><i class="icon-sider-interest"></i>感兴趣</a>
                </li>
                <li data-value="contact">
                    <a href="javascript:;" ka="side_chat"><i class="icon-sider-chat"></i>沟通过</a>
                </li>
                <li data-value="deliver">
                    <a href="javascript:;" ka="side_deliver"><i class="icon-sider-resume"></i>已投递</a>
                </li>
                <li data-value="interview">
                    <a href="javascript:;" ka="side_interview"><i class="icon-sider-interview"></i>面试</a>
                </li>
            </ul>
            <ul class="siderbar-bottom">
                <li>
                    <a class="siderbar-back-top" href="javascript:;" title="返回顶部"></a>
                </li>
                <li>
                    <a class="siderbar-feedback" href="javascript:;">反馈</a>
                </li>
                <li>
                    <a class="siderbar-wechat" href="javascript:;">微信
                        <div class="qrcode-layer "><i></i><img src="https://static.zhipin.com/v2/web/geek/images/wechat-qrcode-2.jpg" alt="">关注BOSS直聘微信服务号</div>
                    </a>
                </li>
                <li>
                    <a class="siderbar" href="https://weibo.com/bosszhipin" target="_blank">微博</a>
                </li>
            </ul>
        </div>
    </div>

    <!--通用注册登录-->
    <div class="sign-wrap sign-wrap-v2" style="display: none">





<!--密码登录-->
<div class="sign-form sign-pwd" style="display:none;">
    <div class="sign-slide-box">
        <a class="logo" href="https://www.zhipin.com/" ka="header-logo">
            <img src="https://static.zhipin.com/v2/web/geek/images/logo.png">
            <div>
                <p>找工作</p >
                <p>我要跟老板谈</p >
            </div>
        </a>
        <ul>
            <li>
                <i class="icon"></i>
                <span>沟通</span>
                <span>在线职位及时沟通</span>
            </li>
            <li>
                <i class="icon"></i>
                <span>任性选</span>
                <span>各大行业职位任你选</span>
            </li>
        </ul>
    </div>
    <div class="sign-content">
        <div class="inner-box">
            <h3 class="title">登录BOSS直聘</h3>
            <form action="/wapi/zppassport/login/account" method="post">
                <div class="sign-tab"><span class="link-signin cur">密码登录</span><span class="link-sms">短信登录</span><span class="link-scan">扫码登录</span></div>
                <div class="tip-error tip-error-form"></div>
                <div class="form-row row-select">
                    <span class="dropdown-select"><i class="icon-select-arrow"></i><em class="text-select">+86</em><input type="hidden" name="regionCode" value="+86"/></span>
                    <span class="ipt-wrap"><i class="icon-sign-phone"></i><input type="tel" class="ipt ipt-phone required" ka="signin-account" placeholder="手机号" name="account"/></span>
                    <div class="dropdown-menu">
<ul>
    <li data-val="+86"><span class="num">+86</span>中国大陆</li>
    <li data-val="+852"><span class="num">+852</span>中国香港</li>
    <li data-val="+853"><span class="num">+853</span>中国澳门</li>
    <li data-val="+886"><span class="num">+886</span>中国台湾</li>
    <li data-val="+1"><span class="num">+1</span>美国</li>
    <li data-val="+81"><span class="num">+81</span>日本</li>
    <li data-val="+44"><span class="num">+44</span>英国</li>
    <li data-val="+82"><span class="num">+82</span>韩国</li>
    <li data-val="+33"><span class="num">+33</span>法国</li>
    <li data-val="+7"><span class="num">+7</span>俄罗斯</li>
    <li data-val="+39"><span class="num">+39</span>意大利</li>
    <li data-val="+65"><span class="num">+65</span>新加坡</li>
    <li data-val="+63"><span class="num">+63</span>菲律宾</li>
    <li data-val="+60"><span class="num">+60</span>马来西亚</li>
    <li data-val="+64"><span class="num">+64</span>新西兰</li>
    <li data-val="+34"><span class="num">+34</span>西班牙</li>
    <li data-val="+95"><span class="num">+95</span>缅甸</li>
    <li data-val="+230"><span class="num">+230</span>毛里求斯</li>
    <li data-val="+263"><span class="num">+263</span>津巴布韦</li>
    <li data-val="+20"><span class="num">+20</span>埃及</li>
    <li data-val="+92"><span class="num">+92</span>巴基斯坦</li>
    <li data-val="+880"><span class="num">+880</span>孟加拉国</li>
</ul>                    </div>
                    <div class="tip-error"></div>
                </div>
                <div class="form-row">
                    <span class="ipt-wrap"><i class="icon-sign-pwd"></i><input type="password" class="ipt ipt-pwd required" ka="signin-password" placeholder="密码" name="password"/></span>
                    <div class="tip-error"></div>
                </div>
                <div class="form-row">
                    <div class="row-code" id="pwdVerrifyCode">
                    </div>
                    <div class="tip-error"></div>
                </div>
                <input type="hidden" name="csessionId">
                <input type="hidden" name="csig">
                <input type="hidden" name="ctoken">
                <input type="hidden" name="cscene" value="nc_login">
                <input type="hidden" value="FFFF0N00000000006DC1" name="cappKey">
                <div class="form-btn">
                    <button type="submit" class="btn">登录</button>
                </div>
                     <div class="text-tip">没有账号 <a href="javascript:;" class="link-signup">立即注册</a></div>
            </form>
        </div>
    </div>
</div>

<!--短信登录-->
<div class="sign-form sign-sms" style="display:none;">
    <div class="sign-slide-box">
        <a class="logo" href="https://www.zhipin.com/" ka="header-logo">
            <img src="https://static.zhipin.com/v2/web/geek/images/logo.png">
            <div>
                <p>找工作</p >
                <p>我要跟老板谈</p >
            </div>
        </a>
        <ul>
            <li>
                <i class="icon"></i>
                <span>沟通</span>
                <span>在线职位及时沟通</span>
            </li>
            <li>
                <i class="icon"></i>
                <span>任性选</span>
                <span>各大行业职位任你选</span>
            </li>
        </ul>
    </div>
    <div class="sign-content">
        <div class="inner-box">
            <h3 class="title">登录BOSS直聘</h3>
            <form action="/wapi/zppassport/login/phone" method="post">
                <div class="sign-tab"><span class="link-signin">密码登录</span><span class="link-sms cur">短信登录</span><span class="link-scan">扫码登录</span></div>
                <div class="tip-error tip-error-form"></div>
                <div class="form-row row-select">
                    <span class="dropdown-select"><i class="icon-select-arrow"></i><em class="text-select">+86</em><input type="hidden" name="regionCode" value="+86"/></span>
                    <span class="ipt-wrap"><i class="icon-sign-phone"></i><input type="tel" class="ipt ipt-phone required" ka="signin-account" placeholder="手机号" name="phone"/></span>
                    <div class="dropdown-menu">
<ul>
    <li data-val="+86"><span class="num">+86</span>中国大陆</li>
    <li data-val="+852"><span class="num">+852</span>中国香港</li>
    <li data-val="+853"><span class="num">+853</span>中国澳门</li>
    <li data-val="+886"><span class="num">+886</span>中国台湾</li>
    <li data-val="+1"><span class="num">+1</span>美国</li>
    <li data-val="+81"><span class="num">+81</span>日本</li>
    <li data-val="+44"><span class="num">+44</span>英国</li>
    <li data-val="+82"><span class="num">+82</span>韩国</li>
    <li data-val="+33"><span class="num">+33</span>法国</li>
    <li data-val="+7"><span class="num">+7</span>俄罗斯</li>
    <li data-val="+39"><span class="num">+39</span>意大利</li>
    <li data-val="+65"><span class="num">+65</span>新加坡</li>
    <li data-val="+63"><span class="num">+63</span>菲律宾</li>
    <li data-val="+60"><span class="num">+60</span>马来西亚</li>
    <li data-val="+64"><span class="num">+64</span>新西兰</li>
    <li data-val="+34"><span class="num">+34</span>西班牙</li>
    <li data-val="+95"><span class="num">+95</span>缅甸</li>
    <li data-val="+230"><span class="num">+230</span>毛里求斯</li>
    <li data-val="+263"><span class="num">+263</span>津巴布韦</li>
    <li data-val="+20"><span class="num">+20</span>埃及</li>
    <li data-val="+92"><span class="num">+92</span>巴基斯坦</li>
    <li data-val="+880"><span class="num">+880</span>孟加拉国</li>
</ul>                    </div>
                    <div class="tip-error"></div>
                </div>
                <div class="form-row">
                    <div class="row-code" id="loginVerrifyCode">
                    </div>
                    <div class="tip-error"></div>
                </div>
                <input type="hidden" name="csessionId">
                <input type="hidden" name="csig">
                <input type="hidden" name="ctoken">
                <input type="hidden" name="cscene" value="nc_login">
                <input type="hidden" value="FFFF0N00000000006DC1" name="cappKey">
                <div class="form-row">
                        <span class="ipt-wrap"><i class="icon-sign-sms"></i><input type="text" class="ipt ipt-sms required" ka="signup-sms" placeholder="短信验证码" name="phoneCode" maxlength="6"/><input
                                type="hidden" name="smsType" value="1"/><button type="button" class="btn btn-sms" data-url="/wapi/zppassport/send/smsCode">发送验证码</button></span>
                    <div class="tip-error"></div>
                </div>
                <div class="form-btn">
                    <button type="submit" class="btn">登录</button>
                </div>
                    <div class="text-tip">没有账号 <a href="javascript:;" class="link-signup">立即注册</a></div>
            </form>
        </div>
    </div>
</div>

<!--扫码登录-->
<div class="sign-form sign-scan" style="display:none;">
    <div class="sign-slide-box">
        <a class="logo" href="https://www.zhipin.com/" ka="header-logo">
            <img src="https://static.zhipin.com/v2/web/geek/images/logo.png">
            <div>
                <p>找工作</p >
                <p>我要跟老板谈</p >
            </div>
        </a>
        <ul>
            <li>
                <i class="icon"></i>
                <span>沟通</span>
                <span>在线职位及时沟通</span>
            </li>
            <li>
                <i class="icon"></i>
                <span>任性选</span>
                <span>各大行业职位任你选</span>
            </li>
        </ul>
    </div>
    <div class="sign-content">
        <div class="inner-box">
            <div class="tip-error tip-error-form"></div>
            <div class="sign-tab"><span class="link-signin">密码登录</span><span class="link-sms">短信登录</span><span class="link-scan cur">扫码登录</span></div>
            <div class="qrcode-box">
                <input type="hidden" class="uuid" value=""/>
                <p><span>使用 BOSS直聘 APP 扫码登录</span><em>扫码帮助</em></p>
                <div class="qrcodeimg-box">
                    <div class="invalid-box">
                        <p>请重新刷新二维码</p>
                        <button class="btn">点击刷新</button>
                    </div>
                    <img src="" data-url="https://www.zhipin.com/qrscan/dispatcher?qrId="/>
                </div>
                <div class="qrcode-tip"><span class="gray">知道了</span>Boss现在也可以使用密码和短信登录了</div>
                <div class="hover-range-left"></div>
                <div class="hover-range-right"></div>
                <div class="sign-scan-help animation">
                    <h4>扫码帮助</h4>
                    <ul class="scan-help-tab">
                        <li class="active">我是BOSS</li>
                        <li>我是牛人</li>
                    </ul>
                    <div class="sub-title">我的&nbsp;&gt;&nbsp;登录网页版</div>
                    <ul class="scan-help-content">
                        <li class="help-boss active"></li>
                        <li class="help-geek"></li>
                    </ul>
                </div>
            </div>
            <div class="text-tip">没有账号 <a href="javascript:;" class="link-signup">立即注册</a></div>
            <div class="login-step-box" >
                <div class="user-photo">
                    <img src="">
                </div>
                <div class="login-step-text">
                    <h3 class="login-step-title">扫描成功</h3>
                    <p class="login-step-detail">请在App端确认登录</p>
                </div>
            </div>
        </div>
    </div>
</div>

<!--注册-->
<div class="sign-form sign-register" style="display:block;">
    <div class="sign-slide-box">
        <a class="logo" href="https://www.zhipin.com/" ka="header-logo">
            <img src="https://static.zhipin.com/v2/web/geek/images/logo.png">
            <div>
                <p>找工作</p >
                <p>我要跟老板谈</p >
            </div>
        </a>
        <ul class="geek-slide ">
            <li>
                <i class="icon"></i>
                <span>沟通</span>
                <span>在线职位及时沟通</span>
            </li>
            <li>
                <i class="icon"></i>
                <span>任性选</span>
                <span>各大行业职位任你选</span>
            </li>
        </ul>
        <ul class="boss-slide hide">
            <li>
                <i class="icon"></i>
                <span>招聘效果好</span>
                <span>与职场牛人在线开聊</span>
            </li>
            <li>
                <i class="icon"></i>
                <span>更多在线牛人</span>
                <span>入职速度快</span>
            </li>
            <li>
                <i class="icon"></i>
                <span>人才匹配度高</span>
                <span>获取更精准的牛人</span>
            </li>
        </ul>
    </div>
    <div class="sign-content">
        <div class="inner-box">
            <h3 class="title">注册BOSS直聘</h3>
            <form action="/wapi/zppassport/user/registered" method="post">
                <input type="hidden" name="act" value="">
                <input type="hidden" name="purpose" value="">
                <h4></h4>
                <div class="tip-error tip-error-form"></div>
                <div class="form-row row-select">
                    <span class="dropdown-select"><i class="icon-select-arrow"></i><em class="text-select">+86</em><input type="hidden" name="regionCode" value="+86"/></span>
                    <span class="ipt-wrap"><i class="icon-sign-phone"></i><input type="tel" class="ipt ipt-phone required" ka="signin-account" placeholder="手机号" name="phone"/></span>
                    <div class="dropdown-menu">
<ul>
    <li data-val="+86"><span class="num">+86</span>中国大陆</li>
    <li data-val="+852"><span class="num">+852</span>中国香港</li>
    <li data-val="+853"><span class="num">+853</span>中国澳门</li>
    <li data-val="+886"><span class="num">+886</span>中国台湾</li>
    <li data-val="+1"><span class="num">+1</span>美国</li>
    <li data-val="+81"><span class="num">+81</span>日本</li>
    <li data-val="+44"><span class="num">+44</span>英国</li>
    <li data-val="+82"><span class="num">+82</span>韩国</li>
    <li data-val="+33"><span class="num">+33</span>法国</li>
    <li data-val="+7"><span class="num">+7</span>俄罗斯</li>
    <li data-val="+39"><span class="num">+39</span>意大利</li>
    <li data-val="+65"><span class="num">+65</span>新加坡</li>
    <li data-val="+63"><span class="num">+63</span>菲律宾</li>
    <li data-val="+60"><span class="num">+60</span>马来西亚</li>
    <li data-val="+64"><span class="num">+64</span>新西兰</li>
    <li data-val="+34"><span class="num">+34</span>西班牙</li>
    <li data-val="+95"><span class="num">+95</span>缅甸</li>
    <li data-val="+230"><span class="num">+230</span>毛里求斯</li>
    <li data-val="+263"><span class="num">+263</span>津巴布韦</li>
    <li data-val="+20"><span class="num">+20</span>埃及</li>
    <li data-val="+92"><span class="num">+92</span>巴基斯坦</li>
    <li data-val="+880"><span class="num">+880</span>孟加拉国</li>
</ul>                    </div>
                    <div class="tip-error"></div>
                </div>
                <div class="form-row">
                    <div class="row-code" id="regVerrifyCode"></div>
                    <div class="tip-error"></div>
                </div>

                <input type="hidden" name="csessionId">
                <input type="hidden" name="csig">
                <input type="hidden" name="ctoken">
                <input type="hidden" name="cscene" value="nc_login">
                <input type="hidden" value="FFFF0N00000000006DC1" name="cappKey">
                <div class="form-row">
                  <span class="ipt-wrap"><i class="icon-sign-sms"></i>
                      <input type="text" class="ipt ipt-sms required" ka="signup-sms" placeholder="短信验证码" name="phoneCode" maxlength="6"/>
                      <input type="hidden" name="smsType" value="2"/>
                      <button type="button" class="btn btn-sms" data-url="/wapi/zppassport/send/smsCode">发送验证码</button>
                  </span>
                    <div class="tip-error"></div>
                </div>
                <div class="form-btn">
                    <button type="submit" class="btn">登录/注册</button>
                </div>
                <div class="text-tip">
                  <div class="tip-error"></div>
                  <input type="checkbox" class="agree-policy" name="policy">同意BOSS直聘<a href="https://about.zhipin.com/agreement?id=registerprotocol" ka="link-privacy1" class="user-agreement" target="_blank">《用户协议》</a>
                  <a href="https://about.zhipin.com/agreement?id=personalinfopro" ka="link-privacy2" class="user-agreement" target="_blank">《个人信息保护政策》</a><br>
                  已有账号<a href="javascript:;" class="link-signin">直接登录</a >
                </div>
            </form>
        </div>
        <div class="btn-switch ewm-switch">
            <div class="switch-tip">微信极速注册<i class="icon-triangle"></i></div>
        </div>
    </div>
</div>

<!--扫码小程序注册-->
<div class="sign-form sign-miniapp" style="display: none;">
    <div class="sign-slide-box">
        <a class="logo" href="https://www.zhipin.com/" ka="header-logo">
            <img src="https://static.zhipin.com/v2/web/geek/images/logo.png">
            <div>
                <p>找工作</p >
                <p>我要跟老板谈</p >
            </div>
        </a>
        <ul>
            <li>
                <i class="icon"></i>
                <span>沟通</span>
                <span>在线职位及时沟通</span>
            </li>
            <li>
                <i class="icon"></i>
                <span>任性选</span>
                <span>各大行业职位任你选</span>
            </li>
        </ul>
    </div>
    <div class="sign-content">
        <div class="inner-box">
            <input type="hidden" class="scene" value="b909d520d4914b4cbe742c4d8aa95f0a"/>
            <h4>微信快速注册</h4>
            <div class="qrcodeimg-box">
                <div class="invalid-box">
                    <p>请重新刷新二维码</p>
                    <button class="btn">点击刷新</button>
                </div>
                <img src="" data-url="https://www.zhipin.com/qrscan/dispatcher?qrId=b909d520d4914b4cbe742c4d8aa95f0a">
            </div>
            <p class="sign-tip">请用微信“扫一扫”扫描上方二维码<br/>进入快捷注册</p>
            <div class="text-tip">已有账号 <a href="javascript:;" class="link-signin">立即登录</a></div>
        </div>
        <div class="btn-switch phone-switch">
            <div class="switch-tip">手机号注册在这里<i class="icon-triangle"></i></div>
        </div>
    </div>
</div>

<!--扫码小程序注册成功-->
<div class="sign-form sign-succ">
    <h3 class="title">注册BOSS直聘</h3>
    <img class="img-succ" src="https://static.zhipin.com/zhipin-geek/v463/web/geek/images/success.png"/>
    <p>注册成功，即将跳转完善流程</p>
    <div class="btn-switch phone-switch"></div>
</div>

<!--注册成功完善简历-->
<div class="sign-form sign-welcome">
    <h3 class="title">欢迎来到BOSS直聘</h3>
    <div class="welcome-box">
        <div class="welcome-text">
            <b>快速完善简历</b>
            <p>做好与Boss对话前的准备吧。</p>
        </div>
        <img src="https://static.zhipin.com/v2/web/geek/images/icon-sign-welcome.png" alt=""/>
        <div class="form-btn"><a ka="registe-complete" href="/web/geek/guide" class="btn">开始完善</a>
        </div>
        <div class="count-down"><em class="num">3s</em> 后自动跳转</div>
    </div>
</div>    </div>
</div>
<script src="https://static.zhipin.com/library/js/lib/jquery-1.12.2.min.js"></script>
<script src="https://static.zhipin.com/zhipin-geek/v463/web/geek/js/main.js"></script>
<input type="hidden" id="page_key_name" value="cpc_job_list" />
<script>
    function get_share_datas_from_html_inapp() {
        var shid = "shdefault",
                urlShid,
                urlSid,
                pk = "pkdefault",
                pp = "ppdefault",
                pkn = (typeof _PK != 'undefined' ? _PK : document.getElementById("page_key_name")),
                getQueryString = function(name) {
                    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"),
                            r = window.location.search.substr(1).match(reg);
                    if (r != null) {
                        return unescape(r[2])
                    }
                    return null;
                };
        urlShid = getQueryString("shid");
        urlSid = getQueryString("sid");
        if (urlShid) {
            shid = urlShid;
        } else if (urlSid) {
            shid = urlSid;
        }
        if (pkn) {
            var pknVal = pkn.value;
            if (pknVal) {
                var pkArray = pknVal.split("|");
                if (pkArray.length == 1) {
                    pk = pkArray[0];
                } else if (pkArray.length >= 2) {
                    pk = pkArray[0];
                    pp = pkArray[1];
                }
            }
        }
        var ret = [];
        ret["shid"] = shid;
        ret["pk"] = pk;
        ret["pp"] = pp;
        return ret;
    }
    function getQueryString(name)
    {
        var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)");
        var r = window.location.search.substr(1).match(reg);
        if(r!=null)return  unescape(r[2]); return null;
    }
</script>
<script>
    var _T = _T || [];
    (function() {
        var script = document.createElement("script");
        script.src = "https://static.zhipin.com/library/js/analytics/ka.zhipin.min.js";
        var s = document.getElementsByTagName("script")[0];
        s.parentNode.insertBefore(script, s);
    })();

    var _hmt = _hmt || [];
    (function() {
        var hm = document.createElement("script");
        hm.src = "//hm.baidu.com/hm.js?194df3105ad7148dcf2b98a91b5e727a";
        var s = document.getElementsByTagName("script")[0];
        s.parentNode.insertBefore(hm, s);
    })();
</script><script type="application/ld+json">
    {
        "@context": "https://zhanzhang.baidu.com/contexts/cambrian.jsonld",
        "@id": "https://www.zhipin.com/c101010100-p100199/?query=golang",
        "title": "「北京golang后端开发招聘」-2021年北京golang后端开发人才招聘信息 - BOSS直聘",
        "description": "BOSS直聘为求职者提供2021年北京golang后端开发招聘信息，百万Boss在线直聘，直接开聊，在线面试，找工作就上BOSS直聘网站或APP，直接与Boss开聊吧！",
        "upDate": "2021-06-01T21:07:23"
    }
</script>
</body>
</html>
`
