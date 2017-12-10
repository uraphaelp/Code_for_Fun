package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {
    var P1, P2, valid, choice []string
    var seed, nonseed string
    var l, index1, index2, i, j int
    P1=[]string{"A1曼联(英)", "B1巴黎(法)", "C1罗马(意)", "D1巴萨(西)", "E1利物浦(英)", "F1曼城(英)", "G1贝西克塔斯(土)", "H1热刺(英)"}
    P2=[]string{"A2巴塞尔(瑞)", "B2拜仁(德)", "C2切尔西(英)", "D2尤文图斯(意)", "E2塞维利亚(西)", "F2矿工(乌)", "G2波尔图(葡)", "H2皇马(西)"}
    l=len(P1)
    //l=1时随机数序列无法产生
    for l>1 {
        para:=rand.New(rand.NewSource(time.Now().UnixNano()))
        index2=para.Intn(l-1)
        //先选非种子队nonseed
        nonseed=P2[index2]
        P2=append(P2[:index2], P2[index2+1:]...) /*删除被取出的队伍*/
        //列出该非种子队所有可能的对手集合valid
        for i=0; i<l; i++ {
            if isLegal(P1[i], nonseed) {
                valid=append(valid, P1[i])
            }
        }
        //抽取可能对手中的1个种子队对手seed
        para=rand.New(rand.NewSource(time.Now().UnixNano()))
        index1=para.Intn(len(valid)-1)
        seed=valid[index1]
        index1=find(P1, seed)
        //输出对阵情况
        fmt.Printf("%s vs %s\n", nonseed, seed)
        fmt.Printf("\n")
        P1=append(P1[:index1], P1[index1+1:]...) /*删除被取出的队伍*/     
        l--
        valid=[]string{}
        //判断剩余的每个非种子队中是否有队伍仅存1支可选种子队
        for i=0; i<l; i++ {
            for j=0; j<l; j++ {
                //如果可选，加入可选集合choice
                if isLegal(P2[i], P1[j]) {
                    choice=append(choice, P1[j])
                }
            }
            //可选集合只有1个队伍，立刻输出对阵
            if len(choice)==1 {
                fmt.Printf("%s vs %s\n", P2[i], choice[0])
                fmt.Printf("\n")
                P2=append(P2[:i], P2[i+1:]...)
                index1=find(P1, choice[0])
                P1=append(P1[:index1], P1[index1+1:]...)
                i--
                l--
            }
            choice=[]string{}    
        }
        //同上：判断剩余的每个种子队中是否有队伍仅存1支可选非种子队
        for i=0; i<l; i++ {
            for j=0; j<l; j++ {
                if isLegal(P1[i], P2[j]) {
                    choice=append(choice, P2[j])
                }
            }
            if len(choice)==1 {
                fmt.Printf("%s vs %s\n", choice[0], P1[i])
                fmt.Printf("\n")
                P1=append(P1[:i], P1[i+1:]...)
                index2=find(P2, choice[0])
                P2=append(P2[:index2], P2[index2+1:]...)
                i--
                l--
            }
            choice=[]string{}    
        }
    }
    //大循环完毕后种子、非种子队集合分别仅存1支队伍
    fmt.Printf("%s vs %s\n", P2[0], P1[0])
}

//判别对阵是否合法
func isLegal(team1, team2 string) bool {
    //不来自同联赛；不来自同小组
    if team1[len(team1)-2]==team2[len(team2)-2] || team1[0]==team2[0] {
        return false
    } else {
        return true
    }
}

//辅助函数：找到被抽取队伍来自原集合中的哪个索引位置，便于删除切片元素
func find(arr []string, s string) int {
    for i:=0; i<len(arr); i++ {
        if s==arr[i] {
            return i
        }
    }
    return 0
}
