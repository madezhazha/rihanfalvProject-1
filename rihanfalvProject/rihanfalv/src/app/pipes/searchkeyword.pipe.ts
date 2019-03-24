import { Pipe, PipeTransform,Injectable } from '@angular/core';
import {DomSanitizer} from '@angular/platform-browser';

@Pipe({
  name: 'keyword'
})
export class SearchkeywordPipe implements PipeTransform {
constructor(private DOM: DomSanitizer) {
  }
 

  transform(val: string, keyword: string[]): any {
    
    if(!val){
      return
    }
    let i=0,len=keyword.length,str=val
    str=str.replace(/<[^>]+>/g,"")//去掉html标签
    for(i;i<len;i++){
      if(keyword[i]==" "){
        continue;
      }
        const Reg = new RegExp(keyword[i], 'gi');    //gi为全部。
      //console.log(Reg)
      str = str.replace(Reg, `<a style="color: #ff2424">${keyword[i]}</a>`)
     }
      const res = str;
      return this.DOM.bypassSecurityTrustHtml(res);
  }
}
