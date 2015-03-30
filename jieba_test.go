package jiebago

import (
	"regexp"
	"testing"
)

var (
	test_contents = []string{
		"这是一个伸手不见五指的黑夜。我叫孙悟空，我爱北京，我爱Python和C++。",
		"我不喜欢日本和服。",
		"雷猴回归人间。",
		"工信处女干事每月经过下属科室都要亲口交代24口交换机等技术性器件的安装工作",
		"我需要廉租房",
		"永和服装饰品有限公司",
		"我爱北京天安门",
		"abc",
		"隐马尔可夫",
		"雷猴是个好网站",
		"“Microsoft”一词由“MICROcomputer（微型计算机）”和“SOFTware（软件）”两部分组成",
		"草泥马和欺实马是今年的流行词汇",
		"伊藤洋华堂总府店",
		"中国科学院计算技术研究所",
		"罗密欧与朱丽叶",
		"我购买了道具和服装",
		"PS: 我觉得开源有一个好处，就是能够敦促自己不断改进，避免敞帚自珍",
		"湖北省石首市",
		"湖北省十堰市",
		"总经理完成了这件事情",
		"电脑修好了",
		"做好了这件事情就一了百了了",
		"人们审美的观点是不同的",
		"我们买了一个美的空调",
		"线程初始化时我们要注意",
		"一个分子是由好多原子组织成的",
		"祝你马到功成",
		"他掉进了无底洞里",
		"中国的首都是北京",
		"孙君意",
		"外交部发言人马朝旭",
		"领导人会议和第四届东亚峰会",
		"在过去的这五年",
		"还需要很长的路要走",
		"60周年首都阅兵",
		"你好人们审美的观点是不同的",
		"买水果然后来世博园",
		"买水果然后去世博园",
		"但是后来我才知道你是对的",
		"存在即合理",
		"的的的的的在的的的的就以和和和",
		"I love你，不以为耻，反以为rong",
		"因",
		"",
		"hello你好人们审美的观点是不同的",
		"很好但主要是基于网页形式",
		"hello你好人们审美的观点是不同的",
		"为什么我不能拥有想要的生活",
		"后来我才",
		"此次来中国是为了",
		"使用了它就可以解决一些问题",
		",使用了它就可以解决一些问题",
		"其实使用了它就可以解决一些问题",
		"好人使用了它就可以解决一些问题",
		"是因为和国家",
		"老年搜索还支持",
		"干脆就把那部蒙人的闲法给废了拉倒！RT @laoshipukong : 27日，全国人大常委会第三次审议侵权责任法草案，删除了有关医疗损害责任“举证倒置”的规定。在医患纠纷中本已处于弱势地位的消费者由此将陷入万劫不复的境地。 ",
		"大",
		"",
		"他说的确实在理",
		"长春市长春节讲话",
		"结婚的和尚未结婚的",
		"结合成分子时",
		"旅游和服务是最好的",
		"这件事情的确是我的错",
		"供大家参考指正",
		"哈尔滨政府公布塌桥原因",
		"我在机场入口处",
		"邢永臣摄影报道",
		"BP神经网络如何训练才能在分类时增加区分度？",
		"南京市长江大桥",
		"应一些使用者的建议，也为了便于利用NiuTrans用于SMT研究",
		"长春市长春药店",
		"邓颖超生前最喜欢的衣服",
		"胡锦涛是热爱世界和平的政治局常委",
		"程序员祝海林和朱会震是在孙健的左面和右面, 范凯在最右面.再往左是李松洪",
		"一次性交多少钱",
		"两块五一套，三块八一斤，四块七一本，五块六一条",
		"小和尚留了一个像大和尚一样的和尚头",
		"我是中华人民共和国公民;我爸爸是共和党党员; 地铁和平门站",
		"张晓梅去人民医院做了个B超然后去买了件T恤",
		"AT&T是一件不错的公司，给你发offer了吗？",
		"C++和c#是什么关系？11+122=133，是吗？PI=3.14159",
		"你认识那个和主席握手的的哥吗？他开一辆黑色的士。",
		"枪杆子中出政权"}

	defaultCutResult = [][]string{[]string{"这是", "一个", "伸手不见五指", "的", "黑夜", "。", "我", "叫", "孙悟空", "，", "我", "爱", "北京", "，", "我", "爱", "Python", "和", "C++", "。"},
		[]string{"我", "不", "喜欢", "日本", "和服", "。"},
		[]string{"雷猴", "回归", "人间", "。"},
		[]string{"工信处", "女干事", "每月", "经过", "下属", "科室", "都", "要", "亲口", "交代", "24", "口", "交换机", "等", "技术性", "器件", "的", "安装", "工作"},
		[]string{"我", "需要", "廉租房"},
		[]string{"永和", "服装", "饰品", "有限公司"},
		[]string{"我", "爱", "北京", "天安门"},
		[]string{"abc"},
		[]string{"隐", "马尔可夫"},
		[]string{"雷猴", "是", "个", "好", "网站"},
		[]string{"“", "Microsoft", "”", "一词", "由", "“", "MICROcomputer", "（", "微型", "计算机", "）", "”", "和", "“", "SOFTware", "（", "软件", "）", "”", "两", "部分", "组成"},
		[]string{"草泥马", "和", "欺实", "马", "是", "今年", "的", "流行", "词汇"},
		[]string{"伊藤", "洋华堂", "总府", "店"},
		[]string{"中国科学院计算技术研究所"},
		[]string{"罗密欧", "与", "朱丽叶"},
		[]string{"我", "购买", "了", "道具", "和", "服装"},
		[]string{"PS", ":", " ", "我", "觉得", "开源", "有", "一个", "好处", "，", "就是", "能够", "敦促", "自己", "不断改进", "，", "避免", "敞帚", "自珍"},
		[]string{"湖北省", "石首市"},
		[]string{"湖北省", "十堰市"},
		[]string{"总经理", "完成", "了", "这件", "事情"},
		[]string{"电脑", "修好", "了"},
		[]string{"做好", "了", "这件", "事情", "就", "一了百了", "了"},
		[]string{"人们", "审美", "的", "观点", "是", "不同", "的"},
		[]string{"我们", "买", "了", "一个", "美的", "空调"},
		[]string{"线程", "初始化", "时", "我们", "要", "注意"},
		[]string{"一个", "分子", "是", "由", "好多", "原子", "组织", "成", "的"},
		[]string{"祝", "你", "马到功成"},
		[]string{"他", "掉", "进", "了", "无底洞", "里"},
		[]string{"中国", "的", "首都", "是", "北京"},
		[]string{"孙君意"},
		[]string{"外交部", "发言人", "马朝旭"},
		[]string{"领导人", "会议", "和", "第四届", "东亚", "峰会"},
		[]string{"在", "过去", "的", "这", "五年"},
		[]string{"还", "需要", "很长", "的", "路", "要", "走"},
		[]string{"60", "周年", "首都", "阅兵"},
		[]string{"你好", "人们", "审美", "的", "观点", "是", "不同", "的"},
		[]string{"买", "水果", "然后", "来", "世博园"},
		[]string{"买", "水果", "然后", "去", "世博园"},
		[]string{"但是", "后来", "我", "才", "知道", "你", "是", "对", "的"},
		[]string{"存在", "即", "合理"},
		[]string{"的", "的", "的", "的", "的", "在", "的", "的", "的", "的", "就", "以", "和", "和", "和"},
		[]string{"I", " ", "love", "你", "，", "不以为耻", "，", "反", "以为", "rong"},
		[]string{"因"},
		[]string{},
		[]string{"hello", "你好", "人们", "审美", "的", "观点", "是", "不同", "的"},
		[]string{"很", "好", "但", "主要", "是", "基于", "网页", "形式"},
		[]string{"hello", "你好", "人们", "审美", "的", "观点", "是", "不同", "的"},
		[]string{"为什么", "我", "不能", "拥有", "想要", "的", "生活"},
		[]string{"后来", "我", "才"},
		[]string{"此次", "来", "中国", "是", "为了"},
		[]string{"使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{",", "使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{"其实", "使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{"好人", "使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{"是因为", "和", "国家"},
		[]string{"老年", "搜索", "还", "支持"},
		[]string{"干脆", "就", "把", "那部", "蒙人", "的", "闲法", "给", "废", "了", "拉倒", "！", "RT", " ", "@", "laoshipukong", " ", ":", " ", "27", "日", "，", "全国人大常委会", "第三次", "审议", "侵权", "责任法", "草案", "，", "删除", "了", "有关", "医疗", "损害", "责任", "“", "举证", "倒置", "”", "的", "规定", "。", "在", "医患", "纠纷", "中本", "已", "处于", "弱势", "地位", "的", "消费者", "由此", "将", "陷入", "万劫不复", "的", "境地", "。", " "},
		[]string{"大"},
		[]string{},
		[]string{"他", "说", "的", "确实", "在理"},
		[]string{"长春", "市长", "春节", "讲话"},
		[]string{"结婚", "的", "和", "尚未", "结婚", "的"},
		[]string{"结合", "成", "分子", "时"},
		[]string{"旅游", "和", "服务", "是", "最好", "的"},
		[]string{"这件", "事情", "的确", "是", "我", "的", "错"},
		[]string{"供", "大家", "参考", "指正"},
		[]string{"哈尔滨", "政府", "公布", "塌桥", "原因"},
		[]string{"我", "在", "机场", "入口处"},
		[]string{"邢永臣", "摄影", "报道"},
		[]string{"BP", "神经网络", "如何", "训练", "才能", "在", "分类", "时", "增加", "区分度", "？"},
		[]string{"南京市", "长江大桥"},
		[]string{"应", "一些", "使用者", "的", "建议", "，", "也", "为了", "便于", "利用", "NiuTrans", "用于", "SMT", "研究"},
		[]string{"长春市", "长春", "药店"},
		[]string{"邓颖超", "生前", "最", "喜欢", "的", "衣服"},
		[]string{"胡锦涛", "是", "热爱", "世界", "和平", "的", "政治局", "常委"},
		[]string{"程序员", "祝", "海林", "和", "朱会震", "是", "在", "孙健", "的", "左面", "和", "右面", ",", " ", "范凯", "在", "最", "右面", ".", "再往", "左", "是", "李松洪"},
		[]string{"一次性", "交", "多少", "钱"},
		[]string{"两块", "五", "一套", "，", "三块", "八", "一斤", "，", "四块", "七", "一本", "，", "五块", "六", "一条"},
		[]string{"小", "和尚", "留", "了", "一个", "像", "大", "和尚", "一样", "的", "和尚头"},
		[]string{"我", "是", "中华人民共和国", "公民", ";", "我", "爸爸", "是", "共和党", "党员", ";", " ", "地铁", "和平门", "站"},
		[]string{"张晓梅", "去", "人民", "医院", "做", "了", "个", "B超", "然后", "去", "买", "了", "件", "T恤"},
		[]string{"AT&T", "是", "一件", "不错", "的", "公司", "，", "给", "你", "发", "offer", "了", "吗", "？"},
		[]string{"C++", "和", "c#", "是", "什么", "关系", "？", "11", "+", "122", "=", "133", "，", "是", "吗", "？", "PI", "=", "3.14159"},
		[]string{"你", "认识", "那个", "和", "主席", "握手", "的", "的哥", "吗", "？", "他开", "一辆", "黑色", "的士", "。"},
		[]string{"枪杆子", "中", "出", "政权"},
	}

	cutAllResult = [][]string{[]string{"这", "是", "一个", "伸手", "伸手不见", "伸手不见五指", "不见", "五指", "的", "黑夜", "", "", "我", "叫", "孙悟空", "悟空", "", "", "我", "爱", "北京", "", "", "我", "爱", "Python", "和", "C++", ""},
		[]string{"我", "不", "喜欢", "日本", "和服", "", ""},
		[]string{"雷猴", "回归", "人间", "", ""},
		[]string{"工信处", "处女", "女干事", "干事", "每月", "月经", "经过", "下属", "科室", "都", "要", "亲口", "口交", "交代", "24", "口交", "交换", "交换机", "换机", "等", "技术", "技术性", "性器", "器件", "的", "安装", "安装工", "装工", "工作"},
		[]string{"我", "需要", "廉租", "廉租房", "租房"},
		[]string{"永和", "和服", "服装", "装饰", "装饰品", "饰品", "有限", "有限公司", "公司"},
		[]string{"我", "爱", "北京", "天安", "天安门"},
		[]string{"abc"},
		[]string{"隐", "马尔可", "马尔可夫", "可夫"},
		[]string{"雷猴", "是", "个", "好", "网站"},
		[]string{"", "Microsoft", "", "一", "词", "由", "", "MICROcomputer", "", "微型", "计算", "计算机", "算机", "", "", "", "和", "", "SOFTware", "", "软件", "", "", "", "两部", "部分", "分组", "组成"},
		[]string{"草泥马", "和", "欺", "实", "马", "是", "今年", "的", "流行", "词汇"},
		[]string{"伊", "藤", "洋华堂", "总府", "店"},
		[]string{"中国", "中国科学院", "中国科学院计算技术研究所", "科学", "科学院", "学院", "计算", "计算技术", "技术", "研究", "研究所"},
		[]string{"罗密欧", "与", "朱丽叶"},
		[]string{"我", "购买", "了", "道具", "和服", "服装"},
		[]string{"PS", "", "", "我", "觉得", "开源", "有", "一个", "好处", "", "", "就是", "能够", "敦促", "自己", "不断", "不断改进", "改进", "", "", "避免", "敞", "帚", "自珍"},
		[]string{"湖北", "湖北省", "石首", "石首市"},
		[]string{"湖北", "湖北省", "十堰", "十堰市"},
		[]string{"总经理", "经理", "理完", "完成", "了", "这件", "事情"},
		[]string{"电脑", "修好", "了"},
		[]string{"做好", "了", "这件", "事情", "就", "一了百了", "了了"},
		[]string{"人们", "审美", "美的", "观点", "是", "不同", "的"},
		[]string{"我们", "买", "了", "一个", "美的", "空调"},
		[]string{"线程", "初始", "初始化", "化时", "我们", "要", "注意"},
		[]string{"一个", "分子", "是", "由", "好多", "原子", "组织", "织成", "的"},
		[]string{"祝", "你", "马到功成"},
		[]string{"他", "掉", "进", "了", "无底", "无底洞", "里"},
		[]string{"中国", "的", "首都", "是", "北京"},
		[]string{"孙", "君", "意"},
		[]string{"外交", "外交部", "部发", "发言", "发言人", "人马", "马朝旭"},
		[]string{"领导", "领导人", "会议", "议和", "第四", "第四届", "四届", "东亚", "峰会"},
		[]string{"在", "过去", "的", "这", "五年"},
		[]string{"还", "需要", "很", "长", "的", "路", "要", "走"},
		[]string{"60", "周年", "首都", "阅兵"},
		[]string{"你好", "好人", "人们", "审美", "美的", "观点", "是", "不同", "的"},
		[]string{"买", "水果", "果然", "然后", "后来", "来世", "世博", "世博园", "博园"},
		[]string{"买", "水果", "果然", "然后", "后去", "去世", "世博", "世博园", "博园"},
		[]string{"但是", "后来", "我", "才", "知道", "你", "是", "对", "的"},
		[]string{"存在", "即", "合理"},
		[]string{"的", "的", "的", "的", "的", "在", "的", "的", "的", "的", "就", "以", "和", "和", "和"},
		[]string{"I", "love", "你", "", "", "不以", "不以为耻", "以为", "耻", "", "", "反", "以为", "rong"},
		[]string{"因"},
		[]string{},
		[]string{"hello", "你好", "好人", "人们", "审美", "美的", "观点", "是", "不同", "的"},
		[]string{"很", "好", "但", "主要", "要是", "基于", "网页", "形式"},
		[]string{"hello", "你好", "好人", "人们", "审美", "美的", "观点", "是", "不同", "的"},
		[]string{"为什么", "什么", "我", "不能", "拥有", "想要", "的", "生活"},
		[]string{"后来", "我", "才"},
		[]string{"此次", "来", "中国", "国是", "为了"},
		[]string{"使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{"", "", "使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{"其实", "使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{"好人", "使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{"是因为", "因为", "和", "国家"},
		[]string{"老年", "搜索", "索还", "支持"},
		[]string{"干脆", "就", "把", "那部", "蒙人", "的", "闲", "法", "给", "废", "了", "拉倒", "", "RT", "", "laoshipukong", "", "", "27", "日", "", "", "全国", "全国人大", "全国人大常委会", "国人", "人大", "人大常委会", "常委", "常委会", "委会", "第三", "第三次", "三次", "审议", "侵权", "权责", "责任", "责任法", "草案", "", "", "删除", "除了", "有关", "医疗", "损害", "责任", "", "", "举证", "倒置", "", "", "的", "规定", "", "", "在", "医患", "纠纷", "中", "本", "已", "处于", "弱势", "地位", "的", "消费", "消费者", "由此", "将", "陷入", "万劫不复", "不复", "的", "境地", "", "", ""},
		[]string{"大"},
		[]string{},
		[]string{"他", "说", "的确", "确实", "实在", "理"},
		[]string{"长春", "长春市", "市长", "长春", "春节", "讲话"},
		[]string{"结婚", "的", "和尚", "尚未", "未结", "结婚", "的"},
		[]string{"结合", "合成", "成分", "分子", "时"},
		[]string{"旅游", "和服", "服务", "是", "最好", "的"},
		[]string{"这件", "事情", "的确", "是", "我", "的", "错"},
		[]string{"供", "大家", "参考", "指正"},
		[]string{"哈尔", "哈尔滨", "政府", "公布", "塌", "桥", "原因"},
		[]string{"我", "在", "机场", "入口", "入口处"},
		[]string{"邢", "永", "臣", "摄影", "报道"},
		[]string{"BP", "神经", "神经网", "神经网络", "网络", "如何", "训练", "才能", "在", "分类", "时", "增加", "加区", "区分", "区分度", "分度", "", ""},
		[]string{"南京", "南京市", "京市", "市长", "长江", "长江大桥", "大桥"},
		[]string{"应", "一些", "使用", "使用者", "用者", "的", "建议", "", "", "也", "为了", "便于", "利用", "NiuTrans", "用于", "SMT", "研究"},
		[]string{"长春", "长春市", "市长", "长春", "春药", "药店"},
		[]string{"邓颖超", "超生", "生前", "最", "喜欢", "的", "衣服"},
		[]string{"胡锦涛", "锦涛", "是", "热爱", "世界", "和平", "的", "政治", "政治局", "常委"},
		[]string{"程序", "程序员", "祝", "海林", "和", "朱", "会", "震", "是", "在", "孙", "健", "的", "左面", "和", "右面", "", "", "", "范", "凯", "在", "最", "右面", "", "", "再往", "左", "是", "李", "松", "洪"},
		[]string{"一次", "一次性", "性交", "多少", "多少钱"},
		[]string{"两块", "五一", "一套", "", "", "三块", "八一", "一斤", "", "", "四块", "七一", "一本", "", "", "五块", "六一", "一条"},
		[]string{"小", "和尚", "留", "了", "一个", "像", "大", "和尚", "一样", "的", "和尚", "和尚头"},
		[]string{"我", "是", "中华", "中华人民", "中华人民共和国", "华人", "人民", "人民共和国", "共和", "共和国", "国公", "公民", "", "", "我", "爸爸", "是", "共和", "共和党", "党员", "", "", "", "地铁", "和平", "和平门", "站"},
		[]string{"张晓梅", "去", "人民", "民医院", "医院", "做", "了", "个", "B", "超然", "然后", "后去", "买", "了", "件", "T", "恤"},
		[]string{"AT", "T", "是", "一件", "不错", "的", "公司", "", "", "给", "你", "发", "offer", "了", "吗", "", ""},
		[]string{"C++", "和", "c#", "是", "什么", "关系", "", "11+122", "133", "", "是", "吗", "", "PI", "3", "14159"},
		[]string{"你", "认识", "那个", "和", "主席", "握手", "的", "的哥", "吗", "", "", "他", "开", "一辆", "黑色", "的士", "", ""},
		[]string{"枪杆", "枪杆子", "杆子", "中出", "政权"},
	}

	defaultCutNoHMMResult = [][]string{[]string{"这", "是", "一个", "伸手不见五指", "的", "黑夜", "。", "我", "叫", "孙悟空", "，", "我", "爱", "北京", "，", "我", "爱", "Python", "和", "C++", "。"},
		[]string{"我", "不", "喜欢", "日本", "和服", "。"},
		[]string{"雷猴", "回归", "人间", "。"},
		[]string{"工信处", "女干事", "每月", "经过", "下属", "科室", "都", "要", "亲口", "交代", "24", "口", "交换机", "等", "技术性", "器件", "的", "安装", "工作"},
		[]string{"我", "需要", "廉租房"},
		[]string{"永和", "服装", "饰品", "有限公司"},
		[]string{"我", "爱", "北京", "天安门"},
		[]string{"abc"},
		[]string{"隐", "马尔可夫"},
		[]string{"雷猴", "是", "个", "好", "网站"},
		[]string{"“", "Microsoft", "”", "一", "词", "由", "“", "MICROcomputer", "（", "微型", "计算机", "）", "”", "和", "“", "SOFTware", "（", "软件", "）", "”", "两", "部分", "组成"},
		[]string{"草泥马", "和", "欺", "实", "马", "是", "今年", "的", "流行", "词汇"},
		[]string{"伊", "藤", "洋华堂", "总府", "店"},
		[]string{"中国科学院计算技术研究所"},
		[]string{"罗密欧", "与", "朱丽叶"},
		[]string{"我", "购买", "了", "道具", "和", "服装"},
		[]string{"PS", ":", " ", "我", "觉得", "开源", "有", "一个", "好处", "，", "就是", "能够", "敦促", "自己", "不断改进", "，", "避免", "敞", "帚", "自珍"},
		[]string{"湖北省", "石首市"},
		[]string{"湖北省", "十堰市"},
		[]string{"总经理", "完成", "了", "这件", "事情"},
		[]string{"电脑", "修好", "了"},
		[]string{"做好", "了", "这件", "事情", "就", "一了百了", "了"},
		[]string{"人们", "审美", "的", "观点", "是", "不同", "的"},
		[]string{"我们", "买", "了", "一个", "美的", "空调"},
		[]string{"线程", "初始化", "时", "我们", "要", "注意"},
		[]string{"一个", "分子", "是", "由", "好多", "原子", "组织", "成", "的"},
		[]string{"祝", "你", "马到功成"},
		[]string{"他", "掉", "进", "了", "无底洞", "里"},
		[]string{"中国", "的", "首都", "是", "北京"},
		[]string{"孙", "君", "意"},
		[]string{"外交部", "发言人", "马朝旭"},
		[]string{"领导人", "会议", "和", "第四届", "东亚", "峰会"},
		[]string{"在", "过去", "的", "这", "五年"},
		[]string{"还", "需要", "很", "长", "的", "路", "要", "走"},
		[]string{"60", "周年", "首都", "阅兵"},
		[]string{"你好", "人们", "审美", "的", "观点", "是", "不同", "的"},
		[]string{"买", "水果", "然后", "来", "世博园"},
		[]string{"买", "水果", "然后", "去", "世博园"},
		[]string{"但是", "后来", "我", "才", "知道", "你", "是", "对", "的"},
		[]string{"存在", "即", "合理"},
		[]string{"的", "的", "的", "的", "的", "在", "的", "的", "的", "的", "就", "以", "和", "和", "和"},
		[]string{"I", " ", "love", "你", "，", "不以为耻", "，", "反", "以为", "rong"},
		[]string{"因"},
		[]string{},
		[]string{"hello", "你好", "人们", "审美", "的", "观点", "是", "不同", "的"},
		[]string{"很", "好", "但", "主要", "是", "基于", "网页", "形式"},
		[]string{"hello", "你好", "人们", "审美", "的", "观点", "是", "不同", "的"},
		[]string{"为什么", "我", "不能", "拥有", "想要", "的", "生活"},
		[]string{"后来", "我", "才"},
		[]string{"此次", "来", "中国", "是", "为了"},
		[]string{"使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{",", "使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{"其实", "使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{"好人", "使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{"是因为", "和", "国家"},
		[]string{"老年", "搜索", "还", "支持"},
		[]string{"干脆", "就", "把", "那", "部", "蒙", "人", "的", "闲", "法", "给", "废", "了", "拉倒", "！", "RT", " ", "@", "laoshipukong", " ", ":", " ", "27", "日", "，", "全国人大常委会", "第三次", "审议", "侵权", "责任法", "草案", "，", "删除", "了", "有关", "医疗", "损害", "责任", "“", "举证", "倒置", "”", "的", "规定", "。", "在", "医患", "纠纷", "中", "本", "已", "处于", "弱势", "地位", "的", "消费者", "由此", "将", "陷入", "万劫不复", "的", "境地", "。", " "},
		[]string{"大"},
		[]string{},
		[]string{"他", "说", "的", "确实", "在", "理"},
		[]string{"长春", "市长", "春节", "讲话"},
		[]string{"结婚", "的", "和", "尚未", "结婚", "的"},
		[]string{"结合", "成", "分子", "时"},
		[]string{"旅游", "和", "服务", "是", "最好", "的"},
		[]string{"这件", "事情", "的确", "是", "我", "的", "错"},
		[]string{"供", "大家", "参考", "指正"},
		[]string{"哈尔滨", "政府", "公布", "塌", "桥", "原因"},
		[]string{"我", "在", "机场", "入口处"},
		[]string{"邢", "永", "臣", "摄影", "报道"},
		[]string{"BP", "神经网络", "如何", "训练", "才能", "在", "分类", "时", "增加", "区分度", "？"},
		[]string{"南京市", "长江大桥"},
		[]string{"应", "一些", "使用者", "的", "建议", "，", "也", "为了", "便于", "利用", "NiuTrans", "用于", "SMT", "研究"},
		[]string{"长春市", "长春", "药店"},
		[]string{"邓颖超", "生前", "最", "喜欢", "的", "衣服"},
		[]string{"胡锦涛", "是", "热爱", "世界", "和平", "的", "政治局", "常委"},
		[]string{"程序员", "祝", "海林", "和", "朱", "会", "震", "是", "在", "孙", "健", "的", "左面", "和", "右面", ",", " ", "范", "凯", "在", "最", "右面", ".", "再", "往", "左", "是", "李", "松", "洪"},
		[]string{"一次性", "交", "多少", "钱"},
		[]string{"两块", "五", "一套", "，", "三块", "八", "一斤", "，", "四块", "七", "一本", "，", "五块", "六", "一条"},
		[]string{"小", "和尚", "留", "了", "一个", "像", "大", "和尚", "一样", "的", "和尚头"},
		[]string{"我", "是", "中华人民共和国", "公民", ";", "我", "爸爸", "是", "共和党", "党员", ";", " ", "地铁", "和平门", "站"},
		[]string{"张晓梅", "去", "人民", "医院", "做", "了", "个", "B超", "然后", "去", "买", "了", "件", "T恤"},
		[]string{"AT&T", "是", "一件", "不错", "的", "公司", "，", "给", "你", "发", "offer", "了", "吗", "？"},
		[]string{"C++", "和", "c#", "是", "什么", "关系", "？", "11", "+", "122", "=", "133", "，", "是", "吗", "？", "PI", "=", "3", ".", "14159"},
		[]string{"你", "认识", "那个", "和", "主席", "握手", "的", "的哥", "吗", "？", "他", "开", "一辆", "黑色", "的士", "。"},
		[]string{"枪杆子", "中", "出", "政权"},
	}

	cutForSearchResult = [][]string{[]string{"这是", "一个", "伸手", "不见", "五指", "伸手不见五指", "的", "黑夜", "。", "我", "叫", "悟空", "孙悟空", "，", "我", "爱", "北京", "，", "我", "爱", "Python", "和", "C++", "。"},
		[]string{"我", "不", "喜欢", "日本", "和服", "。"},
		[]string{"雷猴", "回归", "人间", "。"},
		[]string{"工信处", "干事", "女干事", "每月", "经过", "下属", "科室", "都", "要", "亲口", "交代", "24", "口", "交换", "换机", "交换机", "等", "技术", "技术性", "器件", "的", "安装", "工作"},
		[]string{"我", "需要", "廉租", "租房", "廉租房"},
		[]string{"永和", "服装", "饰品", "有限", "公司", "有限公司"},
		[]string{"我", "爱", "北京", "天安", "天安门"},
		[]string{"abc"},
		[]string{"隐", "可夫", "马尔可", "马尔可夫"},
		[]string{"雷猴", "是", "个", "好", "网站"},
		[]string{"“", "Microsoft", "”", "一词", "由", "“", "MICROcomputer", "（", "微型", "计算", "算机", "计算机", "）", "”", "和", "“", "SOFTware", "（", "软件", "）", "”", "两", "部分", "组成"},
		[]string{"草泥马", "和", "欺实", "马", "是", "今年", "的", "流行", "词汇"},
		[]string{"伊藤", "洋华堂", "总府", "店"},
		[]string{"中国", "科学", "学院", "计算", "技术", "研究", "科学院", "研究所", "中国科学院计算技术研究所"},
		[]string{"罗密欧", "与", "朱丽叶"},
		[]string{"我", "购买", "了", "道具", "和", "服装"},
		[]string{"PS", ":", " ", "我", "觉得", "开源", "有", "一个", "好处", "，", "就是", "能够", "敦促", "自己", "不断", "改进", "不断改进", "，", "避免", "敞帚", "自珍"},
		[]string{"湖北", "湖北省", "石首", "石首市"},
		[]string{"湖北", "湖北省", "十堰", "十堰市"},
		[]string{"经理", "总经理", "完成", "了", "这件", "事情"},
		[]string{"电脑", "修好", "了"},
		[]string{"做好", "了", "这件", "事情", "就", "一了百了", "了"},
		[]string{"人们", "审美", "的", "观点", "是", "不同", "的"},
		[]string{"我们", "买", "了", "一个", "美的", "空调"},
		[]string{"线程", "初始", "初始化", "时", "我们", "要", "注意"},
		[]string{"一个", "分子", "是", "由", "好多", "原子", "组织", "成", "的"},
		[]string{"祝", "你", "马到功成"},
		[]string{"他", "掉", "进", "了", "无底", "无底洞", "里"},
		[]string{"中国", "的", "首都", "是", "北京"},
		[]string{"孙君意"},
		[]string{"外交", "外交部", "发言", "发言人", "马朝旭"},
		[]string{"领导", "领导人", "会议", "和", "第四", "四届", "第四届", "东亚", "峰会"},
		[]string{"在", "过去", "的", "这", "五年"},
		[]string{"还", "需要", "很长", "的", "路", "要", "走"},
		[]string{"60", "周年", "首都", "阅兵"},
		[]string{"你好", "人们", "审美", "的", "观点", "是", "不同", "的"},
		[]string{"买", "水果", "然后", "来", "世博", "博园", "世博园"},
		[]string{"买", "水果", "然后", "去", "世博", "博园", "世博园"},
		[]string{"但是", "后来", "我", "才", "知道", "你", "是", "对", "的"},
		[]string{"存在", "即", "合理"},
		[]string{"的", "的", "的", "的", "的", "在", "的", "的", "的", "的", "就", "以", "和", "和", "和"},
		[]string{"I", " ", "love", "你", "，", "不以", "以为", "不以为耻", "，", "反", "以为", "rong"},
		[]string{"因"},
		[]string{},
		[]string{"hello", "你好", "人们", "审美", "的", "观点", "是", "不同", "的"},
		[]string{"很", "好", "但", "主要", "是", "基于", "网页", "形式"},
		[]string{"hello", "你好", "人们", "审美", "的", "观点", "是", "不同", "的"},
		[]string{"什么", "为什么", "我", "不能", "拥有", "想要", "的", "生活"},
		[]string{"后来", "我", "才"},
		[]string{"此次", "来", "中国", "是", "为了"},
		[]string{"使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{",", "使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{"其实", "使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{"好人", "使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{"因为", "是因为", "和", "国家"},
		[]string{"老年", "搜索", "还", "支持"},
		[]string{"干脆", "就", "把", "那部", "蒙人", "的", "闲法", "给", "废", "了", "拉倒", "！", "RT", " ", "@", "laoshipukong", " ", ":", " ", "27", "日", "，", "全国", "国人", "人大", "常委", "委会", "常委会", "全国人大常委会", "第三", "三次", "第三次", "审议", "侵权", "责任", "责任法", "草案", "，", "删除", "了", "有关", "医疗", "损害", "责任", "“", "举证", "倒置", "”", "的", "规定", "。", "在", "医患", "纠纷", "中本", "已", "处于", "弱势", "地位", "的", "消费", "消费者", "由此", "将", "陷入", "不复", "万劫不复", "的", "境地", "。", " "},
		[]string{"大"},
		[]string{},
		[]string{"他", "说", "的", "确实", "在理"},
		[]string{"长春", "市长", "春节", "讲话"},
		[]string{"结婚", "的", "和", "尚未", "结婚", "的"},
		[]string{"结合", "成", "分子", "时"},
		[]string{"旅游", "和", "服务", "是", "最好", "的"},
		[]string{"这件", "事情", "的确", "是", "我", "的", "错"},
		[]string{"供", "大家", "参考", "指正"},
		[]string{"哈尔", "哈尔滨", "政府", "公布", "塌桥", "原因"},
		[]string{"我", "在", "机场", "入口", "入口处"},
		[]string{"邢永臣", "摄影", "报道"},
		[]string{"BP", "神经", "网络", "神经网", "神经网络", "如何", "训练", "才能", "在", "分类", "时", "增加", "区分", "分度", "区分度", "？"},
		[]string{"南京", "京市", "南京市", "长江", "大桥", "长江大桥"},
		[]string{"应", "一些", "使用", "用者", "使用者", "的", "建议", "，", "也", "为了", "便于", "利用", "NiuTrans", "用于", "SMT", "研究"},
		[]string{"长春", "长春市", "长春", "药店"},
		[]string{"邓颖超", "生前", "最", "喜欢", "的", "衣服"},
		[]string{"锦涛", "胡锦涛", "是", "热爱", "世界", "和平", "的", "政治", "政治局", "常委"},
		[]string{"程序", "程序员", "祝", "海林", "和", "朱会震", "是", "在", "孙健", "的", "左面", "和", "右面", ",", " ", "范凯", "在", "最", "右面", ".", "再往", "左", "是", "李松洪"},
		[]string{"一次", "一次性", "交", "多少", "钱"},
		[]string{"两块", "五", "一套", "，", "三块", "八", "一斤", "，", "四块", "七", "一本", "，", "五块", "六", "一条"},
		[]string{"小", "和尚", "留", "了", "一个", "像", "大", "和尚", "一样", "的", "和尚", "和尚头"},
		[]string{"我", "是", "中华", "华人", "人民", "共和", "共和国", "中华人民共和国", "公民", ";", "我", "爸爸", "是", "共和", "共和党", "党员", ";", " ", "地铁", "和平", "和平门", "站"},
		[]string{"张晓梅", "去", "人民", "医院", "做", "了", "个", "B超", "然后", "去", "买", "了", "件", "T恤"},
		[]string{"AT&T", "是", "一件", "不错", "的", "公司", "，", "给", "你", "发", "offer", "了", "吗", "？"},
		[]string{"C++", "和", "c#", "是", "什么", "关系", "？", "11", "+", "122", "=", "133", "，", "是", "吗", "？", "PI", "=", "3.14159"},
		[]string{"你", "认识", "那个", "和", "主席", "握手", "的", "的哥", "吗", "？", "他开", "一辆", "黑色", "的士", "。"},
		[]string{"枪杆", "杆子", "枪杆子", "中", "出", "政权"},
	}

	cutForSearchNoHMMResult = [][]string{[]string{"这", "是", "一个", "伸手", "不见", "五指", "伸手不见五指", "的", "黑夜", "。", "我", "叫", "悟空", "孙悟空", "，", "我", "爱", "北京", "，", "我", "爱", "Python", "和", "C++", "。"},
		[]string{"我", "不", "喜欢", "日本", "和服", "。"},
		[]string{"雷猴", "回归", "人间", "。"},
		[]string{"工信处", "干事", "女干事", "每月", "经过", "下属", "科室", "都", "要", "亲口", "交代", "24", "口", "交换", "换机", "交换机", "等", "技术", "技术性", "器件", "的", "安装", "工作"},
		[]string{"我", "需要", "廉租", "租房", "廉租房"},
		[]string{"永和", "服装", "饰品", "有限", "公司", "有限公司"},
		[]string{"我", "爱", "北京", "天安", "天安门"},
		[]string{"abc"},
		[]string{"隐", "可夫", "马尔可", "马尔可夫"},
		[]string{"雷猴", "是", "个", "好", "网站"},
		[]string{"“", "Microsoft", "”", "一", "词", "由", "“", "MICROcomputer", "（", "微型", "计算", "算机", "计算机", "）", "”", "和", "“", "SOFTware", "（", "软件", "）", "”", "两", "部分", "组成"},
		[]string{"草泥马", "和", "欺", "实", "马", "是", "今年", "的", "流行", "词汇"},
		[]string{"伊", "藤", "洋华堂", "总府", "店"},
		[]string{"中国", "科学", "学院", "计算", "技术", "研究", "科学院", "研究所", "中国科学院计算技术研究所"},
		[]string{"罗密欧", "与", "朱丽叶"},
		[]string{"我", "购买", "了", "道具", "和", "服装"},
		[]string{"PS", ":", " ", "我", "觉得", "开源", "有", "一个", "好处", "，", "就是", "能够", "敦促", "自己", "不断", "改进", "不断改进", "，", "避免", "敞", "帚", "自珍"},
		[]string{"湖北", "湖北省", "石首", "石首市"},
		[]string{"湖北", "湖北省", "十堰", "十堰市"},
		[]string{"经理", "总经理", "完成", "了", "这件", "事情"},
		[]string{"电脑", "修好", "了"},
		[]string{"做好", "了", "这件", "事情", "就", "一了百了", "了"},
		[]string{"人们", "审美", "的", "观点", "是", "不同", "的"},
		[]string{"我们", "买", "了", "一个", "美的", "空调"},
		[]string{"线程", "初始", "初始化", "时", "我们", "要", "注意"},
		[]string{"一个", "分子", "是", "由", "好多", "原子", "组织", "成", "的"},
		[]string{"祝", "你", "马到功成"},
		[]string{"他", "掉", "进", "了", "无底", "无底洞", "里"},
		[]string{"中国", "的", "首都", "是", "北京"},
		[]string{"孙", "君", "意"},
		[]string{"外交", "外交部", "发言", "发言人", "马朝旭"},
		[]string{"领导", "领导人", "会议", "和", "第四", "四届", "第四届", "东亚", "峰会"},
		[]string{"在", "过去", "的", "这", "五年"},
		[]string{"还", "需要", "很", "长", "的", "路", "要", "走"},
		[]string{"60", "周年", "首都", "阅兵"},
		[]string{"你好", "人们", "审美", "的", "观点", "是", "不同", "的"},
		[]string{"买", "水果", "然后", "来", "世博", "博园", "世博园"},
		[]string{"买", "水果", "然后", "去", "世博", "博园", "世博园"},
		[]string{"但是", "后来", "我", "才", "知道", "你", "是", "对", "的"},
		[]string{"存在", "即", "合理"},
		[]string{"的", "的", "的", "的", "的", "在", "的", "的", "的", "的", "就", "以", "和", "和", "和"},
		[]string{"I", " ", "love", "你", "，", "不以", "以为", "不以为耻", "，", "反", "以为", "rong"},
		[]string{"因"},
		[]string{},
		[]string{"hello", "你好", "人们", "审美", "的", "观点", "是", "不同", "的"},
		[]string{"很", "好", "但", "主要", "是", "基于", "网页", "形式"},
		[]string{"hello", "你好", "人们", "审美", "的", "观点", "是", "不同", "的"},
		[]string{"什么", "为什么", "我", "不能", "拥有", "想要", "的", "生活"},
		[]string{"后来", "我", "才"},
		[]string{"此次", "来", "中国", "是", "为了"},
		[]string{"使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{",", "使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{"其实", "使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{"好人", "使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{"因为", "是因为", "和", "国家"},
		[]string{"老年", "搜索", "还", "支持"},
		[]string{"干脆", "就", "把", "那", "部", "蒙", "人", "的", "闲", "法", "给", "废", "了", "拉倒", "！", "RT", " ", "@", "laoshipukong", " ", ":", " ", "27", "日", "，", "全国", "国人", "人大", "常委", "委会", "常委会", "全国人大常委会", "第三", "三次", "第三次", "审议", "侵权", "责任", "责任法", "草案", "，", "删除", "了", "有关", "医疗", "损害", "责任", "“", "举证", "倒置", "”", "的", "规定", "。", "在", "医患", "纠纷", "中", "本", "已", "处于", "弱势", "地位", "的", "消费", "消费者", "由此", "将", "陷入", "不复", "万劫不复", "的", "境地", "。", " "},
		[]string{"大"},
		[]string{},
		[]string{"他", "说", "的", "确实", "在", "理"},
		[]string{"长春", "市长", "春节", "讲话"},
		[]string{"结婚", "的", "和", "尚未", "结婚", "的"},
		[]string{"结合", "成", "分子", "时"},
		[]string{"旅游", "和", "服务", "是", "最好", "的"},
		[]string{"这件", "事情", "的确", "是", "我", "的", "错"},
		[]string{"供", "大家", "参考", "指正"},
		[]string{"哈尔", "哈尔滨", "政府", "公布", "塌", "桥", "原因"},
		[]string{"我", "在", "机场", "入口", "入口处"},
		[]string{"邢", "永", "臣", "摄影", "报道"},
		[]string{"BP", "神经", "网络", "神经网", "神经网络", "如何", "训练", "才能", "在", "分类", "时", "增加", "区分", "分度", "区分度", "？"},
		[]string{"南京", "京市", "南京市", "长江", "大桥", "长江大桥"},
		[]string{"应", "一些", "使用", "用者", "使用者", "的", "建议", "，", "也", "为了", "便于", "利用", "NiuTrans", "用于", "SMT", "研究"},
		[]string{"长春", "长春市", "长春", "药店"},
		[]string{"邓颖超", "生前", "最", "喜欢", "的", "衣服"},
		[]string{"锦涛", "胡锦涛", "是", "热爱", "世界", "和平", "的", "政治", "政治局", "常委"},
		[]string{"程序", "程序员", "祝", "海林", "和", "朱", "会", "震", "是", "在", "孙", "健", "的", "左面", "和", "右面", ",", " ", "范", "凯", "在", "最", "右面", ".", "再", "往", "左", "是", "李", "松", "洪"},
		[]string{"一次", "一次性", "交", "多少", "钱"},
		[]string{"两块", "五", "一套", "，", "三块", "八", "一斤", "，", "四块", "七", "一本", "，", "五块", "六", "一条"},
		[]string{"小", "和尚", "留", "了", "一个", "像", "大", "和尚", "一样", "的", "和尚", "和尚头"},
		[]string{"我", "是", "中华", "华人", "人民", "共和", "共和国", "中华人民共和国", "公民", ";", "我", "爸爸", "是", "共和", "共和党", "党员", ";", " ", "地铁", "和平", "和平门", "站"},
		[]string{"张晓梅", "去", "人民", "医院", "做", "了", "个", "B超", "然后", "去", "买", "了", "件", "T恤"},
		[]string{"AT&T", "是", "一件", "不错", "的", "公司", "，", "给", "你", "发", "offer", "了", "吗", "？"},
		[]string{"C++", "和", "c#", "是", "什么", "关系", "？", "11", "+", "122", "=", "133", "，", "是", "吗", "？", "PI", "=", "3", ".", "14159"},
		[]string{"你", "认识", "那个", "和", "主席", "握手", "的", "的哥", "吗", "？", "他", "开", "一辆", "黑色", "的士", "。"},
		[]string{"枪杆", "杆子", "枪杆子", "中", "出", "政权"},
	}

	userDictCutResult = [][]string{
		[]string{"这是", "一个", "伸手", "不见", "五指", "的", "黑夜", "。", "我", "叫", "孙悟空", "，", "我", "爱北京", "，", "我", "爱", "Python", "和", "C", "++", "。"},
		[]string{"我", "不", "喜欢", "日本", "和", "服", "。"},
		[]string{"雷猴", "回归人间", "。"},
		[]string{"工信", "处女", "干事", "每", "月", "经过", "下", "属", "科室", "都", "要", "亲口", "交代", "24", "口交换机", "等", "技术性", "器件", "的", "安装", "工作"},
		[]string{"我", "需要", "廉租房"},
		[]string{"永和服", "装饰品", "有", "限公司"},
		[]string{"我", "爱北京", "天安门"},
		[]string{"abc"},
		[]string{"隐马尔", "可夫"},
		[]string{"雷猴", "是", "个", "好", "网站"},
		[]string{"“", "Microsoft", "”", "一词", "由", "“", "MICROcomputer", "（", "微型", "计算机", "）", "”", "和", "“", "SOFTware", "（", "软件", "）", "”", "两部分", "组成"},
		[]string{"草泥", "马", "和", "欺实", "马", "是", "今", "年", "的", "流行", "词汇"},
		[]string{"伊藤洋华堂", "总府", "店"},
		[]string{"中国", "科学院", "计算", "技术", "研究", "所"},
		[]string{"罗密欧", "与", "朱丽叶"},
		[]string{"我购", "买", "了", "道", "具", "和", "服装"},
		[]string{"PS", ":", " ", "我觉", "得", "开源", "有", "一个", "好", "处", "，", "就", "是", "能够", "敦促", "自己", "不断", "改进", "，", "避免", "敞帚", "自珍"},
		[]string{"湖北省", "石首市"},
		[]string{"湖北省", "十堰市"},
		[]string{"总经理", "完成", "了", "这件", "事情"},
		[]string{"电脑", "修好", "了"},
		[]string{"做", "好", "了", "这件", "事情", "就", "一", "了", "百", "了", "了"},
		[]string{"人们", "审美", "的", "观点", "是", "不同", "的"},
		[]string{"我们", "买", "了", "一个", "美", "的", "空调"},
		[]string{"线程", "初始", "化时", "我们", "要", "注意"},
		[]string{"一个", "分子", "是", "由", "好", "多", "原子", "组织成", "的"},
		[]string{"祝", "你", "马到", "功成"},
		[]string{"他", "掉", "进", "了", "无底", "洞里"},
		[]string{"中国", "的", "首", "都", "是", "北京"},
		[]string{"孙君意"},
		[]string{"外交部", "发言人", "马朝旭"},
		[]string{"领导", "人会议", "和", "第四届", "东亚峰", "会"},
		[]string{"在", "过", "去", "的", "这五年"},
		[]string{"还", "需要", "很长", "的", "路", "要", "走"},
		[]string{"60", "周年首", "都", "阅兵"},
		[]string{"你", "好人", "们", "审美", "的", "观点", "是", "不同", "的"},
		[]string{"买水果", "然后", "来", "世博园"},
		[]string{"买水果", "然后", "去", "世博园"},
		[]string{"但", "是", "后", "来", "我", "才", "知道", "你", "是", "对", "的"},
		[]string{"存在", "即", "合理"},
		[]string{"的", "的", "的", "的", "的", "在", "的", "的", "的", "的", "就", "以", "和", "和", "和"},
		[]string{"I", " ", "love", "你", "，", "不以", "为耻", "，", "反以", "为", "rong"},
		[]string{"因"},
		[]string{},
		[]string{"hello", "你", "好人", "们", "审美", "的", "观点", "是", "不同", "的"},
		[]string{"很", "好", "但", "主要", "是", "基于", "网页", "形式"},
		[]string{"hello", "你", "好人", "们", "审美", "的", "观点", "是", "不同", "的"},
		[]string{"为", "什么", "我", "不能", "拥有", "想", "要", "的", "生活"},
		[]string{"后来", "我", "才"},
		[]string{"此次", "来", "中国", "是", "为", "了"},
		[]string{"使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{",", "使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{"其实", "使", "用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{"好人", "使用", "了", "它", "就", "可以", "解决", "一些", "问题"},
		[]string{"是", "因为", "和", "国家"},
		[]string{"老年", "搜索", "还", "支持"},
		[]string{"干脆", "就", "把", "那部", "蒙人", "的", "闲法", "给", "废", "了", "拉", "倒", "！", "RT", " ", "@", "laoshipukong", " ", ":", " ", "27", "日", "，", "全国人", "大常委会", "第三次", "审议", "侵权责", "任法", "草案", "，", "删除", "了", "有", "关医疗", "损害", "责任", "“", "举证", "倒", "置", "”", "的", "规定", "。", "在", "医患", "纠纷", "中本", "已", "处于", "弱势", "地位", "的", "消费者", "由", "此", "将", "陷入", "万劫", "不复", "的", "境地", "。", " "},
		[]string{"大"},
		[]string{},
		[]string{"他", "说", "的", "确实", "在", "理"},
		[]string{"长春市", "长春节", "讲话"},
		[]string{"结婚", "的", "和", "尚未", "结婚", "的"},
		[]string{"结合成", "分子", "时"},
		[]string{"旅游", "和", "服务", "是", "最", "好", "的"},
		[]string{"这件", "事情", "的", "确是", "我", "的", "错"},
		[]string{"供大家", "参考", "指正"},
		[]string{"哈尔滨", "政府", "公布塌桥", "原因"},
		[]string{"我", "在", "机场", "入口", "处"},
		[]string{"邢永臣", "摄影", "报道"},
		[]string{"BP", "神经", "网络", "如何", "训练", "才", "能", "在", "分类", "时", "增加区", "分度", "？"},
		[]string{"南京市", "长江大桥"},
		[]string{"应一些", "使", "用者", "的", "建议", "，", "也", "为", "了", "便", "于", "利用", "NiuTrans", "用于", "SMT", "研究"},
		[]string{"长春市", "长春药店"},
		[]string{"邓颖", "超生", "前", "最", "喜欢", "的", "衣服"},
		[]string{"胡锦涛", "是", "热爱世界", "和", "平", "的", "政治局", "常委"},
		[]string{"程序员", "祝海林", "和", "朱会震", "是", "在", "孙健", "的", "左面", "和", "右面", ",", " ", "范凯", "在", "最右面", ".", "再往", "左", "是", "李松洪"},
		[]string{"一次性", "交多少", "钱"},
		[]string{"两块", "五一套", "，", "三块", "八一斤", "，", "四块", "七", "一本", "，", "五块", "六", "一条"},
		[]string{"小", "和", "尚留", "了", "一个", "像", "大", "和", "尚", "一样", "的", "和", "尚头"},
		[]string{"我", "是", "中华人民共和国", "公民", ";", "我", "爸爸", "是", "共", "和", "党", "党员", ";", " ", "地铁", "和", "平门", "站"},
		[]string{"张晓梅", "去", "人民医院", "做", "了", "个", "B", "超然", "后", "去", "买", "了", "件", "T", "恤"},
		[]string{"AT", "&", "T", "是", "一件", "不错", "的", "公司", "，", "给", "你", "发", "offer", "了", "吗", "？"},
		[]string{"C", "++", "和", "c", "#", "是", "什么", "关系", "？", "11", "+", "122", "=", "133", "，", "是", "吗", "？", "PI", "=", "3.14159"},
		[]string{"你", "认识", "那个", "和", "主席握", "手", "的", "的", "哥", "吗", "？", "他开", "一辆", "黑色", "的", "士", "。"},
		[]string{"枪杆子", "中", "出政权"},
	}
)

func chanToArray(ch chan string) []string {
	result := make([]string, 0)
	for word := range ch {
		result = append(result, word)
	}
	return result
}

func TestCutDAG(t *testing.T) {
	j, _ := Open("dict.txt")

	result := chanToArray(j.cutDAG("BP神经网络如何训练才能在分类时增加区分度？"))
	if len(result) != 11 {
		t.Fatal(result)
	}
}

func TestCutDAGNoHmm(t *testing.T) {
	j, _ := Open("dict.txt")

	result := chanToArray(j.cutDAGNoHMM("BP神经网络如何训练才能在分类时增加区分度？"))
	if len(result) != 11 {
		t.Fatal(result)
	}
}

func TestRegexpSplit(t *testing.T) {
	result := chanToArray(RegexpSplit(regexp.MustCompile(`\p{Han}+`),
		"BP神经网络如何训练才能在分类时增加区分度？"))
	if len(result) != 3 {
		t.Fatal(result)
	}
	result = chanToArray(RegexpSplit(regexp.MustCompile(`([\p{Han}#]+)`),
		",BP神经网络如何训练才能在分类时#增加区分度？"))
	if len(result) != 3 {
		t.Fatal(result)
	}
}

func TestDefaultCut(t *testing.T) {
	j, _ := Open("dict.txt")

	var result []string
	for index, content := range test_contents {
		result = chanToArray(j.Cut(content, true))
		if len(result) != len(defaultCutResult[index]) {
			t.Fatalf("default cut for %s length should be %d not %d\n",
				content, len(defaultCutResult[index]), len(result))
		}
		for i, r := range result {
			if r != defaultCutResult[index][i] {
				t.Fatal(r)
			}
		}
	}
}

func TestCutAll(t *testing.T) {
	j, _ := Open("dict.txt")

	var result []string
	for index, content := range test_contents {
		result = chanToArray(j.CutAll(content))
		if len(result) != len(cutAllResult[index]) {
			t.Fatalf("cut all for %s length should be %d not %d\n",
				content, len(cutAllResult[index]), len(result))
		}
		for i, c := range result {
			if c != cutAllResult[index][i] {
				t.Fatal(c)
			}
		}
	}
}

func TestDefaultCutNoHMM(t *testing.T) {
	j, _ := Open("dict.txt")

	var result []string
	for index, content := range test_contents {
		result = chanToArray(j.Cut(content, false))
		if len(result) != len(defaultCutNoHMMResult[index]) {
			t.Fatalf("default cut no hmm for %s length should be %d not %d\n",
				content, len(defaultCutNoHMMResult[index]), len(result))
		}
		for i, c := range result {
			if c != defaultCutNoHMMResult[index][i] {
				t.Fatal(c)
			}
		}
	}
}

func TestCutForSearch(t *testing.T) {
	j, _ := Open("dict.txt")

	var result []string
	for index, content := range test_contents {
		result = chanToArray(j.CutForSearch(content, true))
		if len(result) != len(cutForSearchResult[index]) {
			t.Fatalf("cut for search for %s length should be %d not %d\n",
				content, len(cutForSearchResult[index]), len(result))
		}
		for i, c := range result {
			if c != cutForSearchResult[index][i] {
				t.Fatal(c)
			}
		}
	}
	for index, content := range test_contents {
		result = chanToArray(j.CutForSearch(content, false))
		if len(result) != len(cutForSearchNoHMMResult[index]) {
			t.Fatalf("cut for search no hmm for %s length should be %d not %d\n",
				content, len(cutForSearchNoHMMResult[index]), len(result))
		}
		for i, c := range result {
			if c != cutForSearchNoHMMResult[index][i] {
				t.Fatal(c)
			}
		}
	}
}

func TestSetdictionary(t *testing.T) {
	var result []string
	j, _ := Open("foobar.txt")
	for index, content := range test_contents {
		result = chanToArray(j.Cut(content, true))
		if len(result) != len(userDictCutResult[index]) {
			t.Fatalf("default cut with user dictionary for %s length should be %d not %d\n",
				content, len(userDictCutResult[index]), len(result))
		}
		for i, c := range result {
			if c != userDictCutResult[index][i] {
				t.Fatal(c)
			}
		}
	}
}

func TestLoadUserDict(t *testing.T) {
	j, _ := Open("dict.txt")
	j.LoadUserDict("userdict.txt")

	sentence := "李小福是创新办主任也是云计算方面的专家; 什么是八一双鹿例如我输入一个带“韩玉赏鉴”的标题，在自定义词库中也增加了此词为N类型"
	result := []string{"李小福", "是", "创新办", "主任", "也", "是", "云计算", "方面", "的", "专家", ";", " ", "什么", "是", "八一双鹿", "例如", "我", "输入", "一个", "带", "“", "韩玉赏鉴", "”", "的", "标题", "，", "在", "自定义词", "库中", "也", "增加", "了", "此", "词为", "N", "类型"}

	words := chanToArray(j.Cut(sentence, true))
	if len(words) != len(result) {
		t.Fatal(len(words))
	}
	for index, word := range words {
		if word != result[index] {
			t.Fatal(word)
		}
	}

	sentence = "easy_install is great"
	result = []string{"easy_install", " ", "is", " ", "great"}
	words = chanToArray(j.Cut(sentence, true))
	if len(words) != len(result) {
		t.Fatal(len(words))
	}
	for index, word := range words {
		if word != result[index] {
			t.Fatal(word)
		}
	}

	sentence = "python 的正则表达式是好用的"
	result = []string{"python", " ", "的", "正则表达式", "是", "好用", "的"}
	words = chanToArray(j.Cut(sentence, true))
	if len(words) != len(result) {
		t.Fatal(words)
		t.Fatal(result)
	}
	for index, word := range words {
		if word != result[index] {
			t.Fatal(word)
		}
	}
}
