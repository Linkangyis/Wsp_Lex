package token

import(
  "strconv"
  "strings"
  "io/ioutil"
)

func token_str_if(str string)(string){
    var maps = map[string]string{}
    maps["\n"] = "T_KH";
    
    if v, ok := maps[str]; ok {
        return v
    }else{
        return str
    }
}

func token(str string,line int)([4]string){
    var tokens = map[string]int{}
    tokens = token_map()
    
    var tokens_name = map[string]string{}
    tokens_name = token_text_map()
    
    if v, ok := tokens[str]; ok {
        return [4]string{strconv.Itoa(v),token_str_if(str),tokens_name[str],strconv.Itoa(line)}
    }else{
        return [4]string{strconv.Itoa(0),str,"T_TEXT",strconv.Itoa(line)}
    }
}

func Wsp_Lexical(file string)(map[int][4]string){
    data, _ := ioutil.ReadFile(file)
    code :=string(data)+"\n"
    num := 0;
    num_TEXTS :=0
    lock_TEXT := 0
    line :=1
    var return_map = map[int][4]string{}
    code_wsp := strings.Split(code,"")
    code_len := len(code_wsp)-1
    for i:=0;i<=code_len;i++{
        tk := token(string(code_wsp[i][len(code_wsp[i])-1]),line);
        //fmt.Println(tk)
        if tk[0] == strconv.Itoa(0){
            if i!=len(code_wsp)-1 {
                tmp := code_wsp[i+1]
                code_wsp[i+1]=code_wsp[i]+code_wsp[i+1]
                num_TEXTS++
                if token(string(code_wsp[i+1][len(code_wsp[i+1])-1]),line)[0] != strconv.Itoa(0){
                    code_wsp[i+1] = tmp
                    for z:=i-num_TEXTS+1;z<=i-1;z++{
                        code_wsp[z]=""
                    }
                    num_TEXTS = 0
                    lock_TEXT = 1
                }
            }
        }else{
            code_wsp[i]=string(code_wsp[i][len(code_wsp[i])-1])
            tk[1]=code_wsp[i]
        }
        if lock_TEXT==1 {
            return_map[num]=token(code_wsp[i],line)
            num++
            lock_TEXT = 0
        }else if token(string(code_wsp[i][len(code_wsp[i])-1]),line)[0] != strconv.Itoa(0){
            return_map[num]=token(code_wsp[i],line)
            num++
        }
        if code_wsp[i]=="\n"{
            line++
        }
    }
    return return_map
}

