import { Component, OnInit } from '@angular/core';
import { ApiSerivice } from '../../../services/apiservice';
import { APIResponse, Nowpage, Allpage } from '../../../services/apiresponse';
import { Router } from '@angular/router';

@Component({
  selector: 'app-legal',
  templateUrl: './legal.component.html',
  styleUrls: ['./legal.component.css']
})
export class LegalComponent implements OnInit {

  constructor(public router: Router, private api: ApiSerivice ) { }

  public types: APIResponse[];
  public category: string;
  public categoryb: string;
  public judge: boolean;
  public Nowcountry: string;   // 当前模块 日/韩
  public pages = 'ture';
  public nowpage: number;
  public paging: number;
  public pagings: any = [0, 0, 0, 0, 0, 0, 0, 0];
  public ifclik: boolean;

  public type: APIResponse = {
    legaltype:    '',
  };

  public page: Nowpage = {
    nowpage:    1,
    kind:      '',
  };

  public allpage: Allpage = {
    allpage:    0,
  };

  ngOnInit() {
    this.judge = false;
    this.ifclik = false;
    this.Postpage(1, 'type');        // 首页页码
    this.getArticles();      // 获取日韩
  }

  public Getpage() {         // 获取总条数
    this.api.legalpage2().subscribe(response => {
      this.allpage = response;
      this.Paging();
    });
  }

  public Postpage(nowpage: number, kind: string) {        // 将当前页数传递会服务器
    this.api.legalpage(nowpage, kind).subscribe(() => {
      this.Getpage();                // 获取总条数
      this.Gettype();    // 获取法条
    });
  }

  public getArticles() {                     // 页面初始化时获取日韩
    const country = localStorage.getItem('JapanOrKorea');
    if (country === '日') {
      this.Nowcountry = 'Japan';
    } else {
      this.Nowcountry = 'Korea';
    }
    this.api.country(this.Nowcountry).subscribe(() => {
      this.Postpage(1, 'type');        // 首页页码
      this.Gettype();    // 获取总页数和法条
    });
  }

  getJapanKorea(isJapan: boolean) {          // 点击按钮时获取日韩
    if (isJapan) {
      this.Nowcountry = 'Japan';
    } else {
      this.Nowcountry = 'Korea';
    }
    this.api.country(this.Nowcountry).subscribe(() => {
      this.Postpage(1, 'type');        // 首页页码
      this.Gettype();    // 获取总页数和法条
    });
  }

  public showAside() {
    this.ifclik   = true;
    // const asideDom: any = document.getElementById('aside');  // 调出侧拉栏
    // asideDom.style.transform = 'translate(0,0)';
   }

   public hideAside() {
     const asideDom: any = document.getElementById('aside');  // 隐藏侧拉栏
     asideDom.style.transform = 'translate(100%,0)';
     const that = this;
     // tslint:disable-next-line:only-arrow-functions
     setTimeout( function() { that.ifclik = false; } , 500);
   }

    public Gettype() {           // 获取法律条文
     this.api.legaltype().subscribe(response => {
       this.types = response;
     });
   }

   public posttype(legaltype: string) {            // 传递选中的标题回服务器
     this.api.legaltype2(legaltype).subscribe(() => {
       this.title();
     });
   }
     public title() {
     this.router.navigate(['article']);
   }

   public choose() {                         // 法律条文的分类显示
    if (this.category === this.categoryb) { // 判断是否选中
       this.judge = !this.judge;
       if (this.judge !== true) {
         this.category = '';
         this.Postpage(1, 'type');        // 首页页码
         this.Gettype();
       } else {
        this.Postpage(1, 'typec');        // 首页页码
        const that = this;
        // tslint:disable-next-line:only-arrow-functions
        setTimeout( function() {
         that.api.legallabel2().subscribe(response => {
         that.types = response;
          }); } , 400);
       }
     } else {
       this.judge = true;
       this.categoryb = this.category;
       this.api.legallabel(this.category).subscribe(() => {
        this.Postpage(1, 'typec');        // 首页页码
       });
       const that = this;
       // tslint:disable-next-line:only-arrow-functions
       setTimeout( function() {
         that.api.legallabel2().subscribe(response => {
         that.types = response; }); } , 300);
     }
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
  if (this.judge !== true) {
    this.Postpage(now, 'type');
    const that = this;
    // tslint:disable-next-line:only-arrow-functions
    setTimeout( function() {
      that.Gettype();
    } , 400);   // 获取法条
  } else {
    this.Postpage(now, 'typec');
    const that2 = this;
    // tslint:disable-next-line:only-arrow-functions
    setTimeout( function() {
      that2.api.legallabel2().subscribe(response => {
        that2.types = response; });
    } , 400);   // 获取法条
  }
// alert('测试弹窗');
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
