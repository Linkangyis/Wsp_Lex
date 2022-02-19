package token

func token_map()(map[string]int){
    var maps = map[string]int{}
    maps["("] = 1;
    maps[")"] = 1;
    maps["{"] = 2;
    maps["}"] = 2;
    maps["["] = 3;
    maps["]"] = 3;
    maps["\n"] = 5;
    maps[" "] = 6;
    maps["="] = 7;
    maps[">"] = 8;
    maps["<"] = 9;
    maps["function"] = 10;
    maps["for"] = 11;
    maps["print"] = 12;
    maps["curl_get"] = 13;
    maps["curl_post"] = 14;
    maps["url_set"] = 15;
    maps["class"] = 20;
    maps["public"] = 21;
    maps["new"] = 21;
    maps["-"] = 22;
    maps["eval"] = 23;
    maps["len"] = 24;
    maps["if"] = 25;
    maps["else"] = 26;
    return maps;
}
func token_text_map()(map[string]string){
    var maps = map[string]string{}
    maps["("] = "D_X_KUO";
    maps[")"] = "D_X_KUO";
    maps["{"] = "D_H_KUO";
    maps["}"] = "D_H_KUO";
    maps["["] = "D_Z_KUO";
    maps["]"] = "D_Z_KUO";
    maps["\n"] = "T_KH";
    maps[" "] = "T_K";
    maps["="] = "Y_DE";
    maps[">"] = "Y_DA - CLASS_START_B";
    maps["<"] = "Y_XA";
    maps["function"] = "H_FUNCTION";
    maps["for"] = "H_FOR";
    maps["print"] = "H_PRINT";
    maps["curl_get"] = "CURL_HTTP_GET";
    maps["curl_post"] = "CURL_HTTP_POST";
    maps["url_set"] = "HTTP_URL_SET";
    maps["class"] = "CLASS_SET";
    maps["public"] = "CLASS_PUBLIC";
    maps["new"] = "CLASS_NEW";
    maps["-"] = "CLASS_START_A";
    maps["eval"] = "H_EVAL";
    maps["len"] = "H_LEN";
    maps["if"] = "H_IF";
    maps["else"] = "H_IF_ELSE";
    return maps;
}
