import { Component, OnInit } from '@angular/core';
// import { ApiSerivice } from 'src/app/apiservice';
// import { APIResponse } from 'src/app/apiresponse';
import { ApiSerivice } from '../../../apiservice';
import { APIResponse } from '../../../apiresponse';
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

  public type: APIResponse = {
    legaltype:    '',
  };

  ngOnInit() {
    this.Gettype();
    this.judge = false;
  }
  public showAside() {
    const asideDom: any = document.getElementById('aside');  // 调出侧拉栏
    asideDom.style.transform = 'translate(0,0)';
   }

   public hideAside() {
     const asideDom: any = document.getElementById('aside');  //隐藏侧拉栏
     asideDom.style.transform = 'translate(100%,0)';
   }

    public Gettype() {           // 获取法律条文
     this.api.legaltype().subscribe(response => {
       this.types = response;
     });
   }

   public posttype(legaltype: string) {            // 传递选中的标题回服务器
     this.api.legaltype2(legaltype).subscribe();
     const that = this;
   setTimeout( function () { that.title(legaltype); } , 200);    // 延时触发，给服务器反应留时间
   }

   public title(legaltype: string) {
     this.router.navigate(['article'], {
       queryParams: {legaltype: legaltype}
       });
   }

   public choose() {                         // 法律条文的分类显示
     if (this.category === this.categoryb) {
       this.judge = !this.judge;
       this.api.legallabel2().subscribe(response => {
         this.types = response; });
       if (this.judge !== true) {
         this.category = '';
         this.Gettype();
       }
     } else {
       this.judge = true;
       this.categoryb = this.category;
       this.api.legallabel(this.category).subscribe();
       const that = this;
       setTimeout( function () {
         that.api.legallabel2().subscribe(response => {
         that.types = response; }); } , 300);
     }
   }
 }
