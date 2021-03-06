import { Component, OnInit } from '@angular/core';
import { ApiSerivice } from '../../../services/apiservice';
import { Router } from '@angular/router';

@Component({
  selector: 'app-content',
  templateUrl: './content.component.html',
  styleUrls: ['./content.component.css']
})
export class ContentComponent implements OnInit {

  constructor(public router: Router, private api: ApiSerivice) { }
  public Nowcountry: string;   // 当前模块 日/韩

  public content: any = {
    legaltype:   '',
    legaltitle:  '',
    legalcontent: ''
  };


  ngOnInit() {
    this.Getcontent();
  }

   public Getcontent() {           // 获取法律条文
    this.api.legalcontent().subscribe(response => {
      this.content = response;
    });
    }

    public back() {       // 返回上一页面
      this.router.navigate(['/legal']);
  }

  public getJapanKorea(isJapan: boolean) {
    if (isJapan) {
      this.back();
    } else {
      this.back();
    }
  }

}
