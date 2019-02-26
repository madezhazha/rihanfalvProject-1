import { Pipe, PipeTransform,Injectable } from '@angular/core';
import {DomSanitizer} from '@angular/platform-browser';

@Pipe({
  name: 'keyword'
})
export class SearchkeywordPipe implements PipeTransform {
constructor(private DOM: DomSanitizer) {
  }
 
 
  
  transform(val: string, keyword: string): any {
    const Reg = new RegExp(keyword, 'gi');
    if (val) {
      const res = val.replace(Reg, `<a style="color: #ff2424;">${keyword}</a>`);
      return this.DOM.bypassSecurityTrustHtml(res);
    }
  }

}
