import { Component, OnInit } from '@angular/core';
import { APIResponse2, Nowpage, Allpage } from '../../../services/apiresponse';
import { ApiSerivice } from '../../../services/apiservice';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-article',
  templateUrl: './article.component.html',
  styleUrls: ['./article.component.css']
})
export class ArticleComponent implements OnInit {

  constructor(public router: Router, private api: ApiSerivice) { }

  public legaltype: string;
  public titles: APIResponse2[];
  public Nowcountry: string;   // 当前模块 日/韩
  public pages = 'ture';
  public nowpage: number;
  public paging: number;
  public pagings: any = [0, 0, 0, 0, 0, 0, 0, 0];

  public page: Nowpage = {
    nowpage:    1,
    kind:      '',
  };

  public allpage: Allpage = {
    allpage:    0,
  };

  public title1: APIResponse2 = {
    legaltitle:    '',
    legaltype:     '',
  };
  public title2: APIResponse2 = {
    legaltitle:    '',
    legaltype:     '',
  };

  ngOnInit() {
    // this.Gettitle()
    this.Postpage(1, 'title');        // 首页页码
    const that = this;
    // tslint:disable-next-line:only-arrow-functions
    setTimeout( function() { that.Getpage(); that.Gettitle(); } , 300);   // 获取总页数和法条
  }

  public Getpage() {         // 获取总条数
    this.api.legalpage2().subscribe(response => {
      this.allpage = response;
      this.Paging();
    });
  }

  public Gettitle() {           // 获取信息
    this.api.legaltitle().subscribe(response => {
      this.titles = response;
      this.legaltype = this.titles[0].legaltype;
    });
  }

  getJapanKorea(isJapan: boolean) {
    if (isJapan) {
      this.back();
    } else {
      this.back();
    }
  }

  public Postpage(nowpage: number, kind: string) {        // 将当前页数传递会服务器
    this.api.legalpage(nowpage, kind).subscribe();
  }

  public posttype(legaltitle: string) {            // 传递标题回服务器
    this.api.legaltitle2(legaltitle).subscribe();
    const that = this;
    // tslint:disable-next-line:only-arrow-functions
    setTimeout( function() { that.content(); } , 200);   // 延时触发，给服务器留反应时间
}
  public content() {
    this.router.navigate(['/content']);
  }

  public back() {    // 返回上一页面
    this.router.navigate(['/legal']);
  }

  public Paging() {
    this.pagings = [0, 0, 0, 0, 0, 0, 0, 0];
    let p = this.page.nowpage;
    const n = this.allpage.allpage / 10;
    if (n > 7) {
     for ( let i = 0; i < 8; i++) {
       if (p > n) {
         this.pagings[i] = 0;
       } else {
       this.pagings[i] = p; }
       p++;
     }
   } else {
     for ( let i = 0; i < n; i++) {
       this.pagings[i] = i + 1;
       p++;
     }
   }
  }

  public Nowpage( now: number) {   // 页面跳转
    this.page.nowpage = now;
    this.Postpage(now, 'title');
    const that = this;
    // tslint:disable-next-line:only-arrow-functions
    setTimeout( function() { that.Gettitle(); } , 400);   // 获取总页数和法条
 }

 public addpage() {
  if (this.page.nowpage < this.allpage.allpage / 1) {
  this.Nowpage(++this.page.nowpage); }
 }

 public reducepage() {
  if (this.page.nowpage > 1) {
  this.Nowpage(--this.page.nowpage); }
}
 }
