import { Component, OnInit } from '@angular/core';
import { APIResponse2 } from 'src/app/apiresponse';
import { ApiSerivice } from 'src/app/apiservice';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-article',
  templateUrl: './article.component.html',
  styleUrls: ['./article.component.css']
})
export class ArticleComponent implements OnInit {

  constructor(private activatedRoute: ActivatedRoute, public router: Router, private api: ApiSerivice) { }

  public legaltype: string;
  public titles: APIResponse2[];

  public title1: APIResponse2 = {
    legaltitle:    '',
  };

  ngOnInit() {
    this.Gettitle();
    this.activatedRoute.queryParams.subscribe(queryParams => {
      this.legaltype = queryParams.legaltype;
  });
  }

  public Gettitle() {           // 获取信息
    this.api.legaltitle().subscribe(response => {
      this.titles = response;
    });
  }

  public posttype(legaltitle: string) {            // 传递标题回服务器
    this.api.legaltitle2(legaltitle).subscribe();
    const that = this;
  setTimeout( function () { that.content(); } , 200);   // 延时触发，给服务器留反应时间
}
  public content() {
    this.router.navigate(['/content']);
  }

  public back() {    // 返回上一页面
    this.router.navigate(['/legal']);
  }
}
