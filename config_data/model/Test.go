     
package config
type Test struct {
    //名称
        Name string;
        
        //描述
        Description string;
        
        //备注
        Comment string;
        
        //图标
        Icon string;
        
        //最大数量
        MaxNumber int;
        
        //能否使用
        CanUse bool;
        
        //能否显示在背包
        CanShowInBag bool;
        
        //是否礼包
        IsPackage bool;
        
        //类型1（0-无，1-英雄碎片，2-强化材料，3-工坊制品，4-基础材料，5-礼包）
        Type1 int;
        
        //类型2
        Type2 int;
        
        //类型3
        Type3 int;
        
        
	
}