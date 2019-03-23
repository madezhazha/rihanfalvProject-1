import { Pipe, PipeTransform,Injectable } from '@angular/core';
import {DomSanitizer} from '@angular/platform-browser';

@Pipe({
  name: 'keyword'
})
export class SearchkeywordPipe implements PipeTransform {
constructor(private DOM: DomSanitizer) {
  }
 
 
  
  /*transform(val: string, keyword: string): any {
    const Reg = new RegExp(keyword, 'gi');
    if (val) {
      var str
      str = val.replace(Reg, `<a style="color: #ff2424;">${keyword}</a>`)
      const res = str;
      return this.DOM.bypassSecurityTrustHtml(res);
    }
  }*/
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
