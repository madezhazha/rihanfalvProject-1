import { Component, OnInit } from '@angular/core';
import { ApiSerivice } from 'src/app/apiservice';
import { Router } from '@angular/router';

@Component({
  selector: 'app-content',
  templateUrl: './content.component.html',
  styleUrls: ['./content.component.css']
})
export class ContentComponent implements OnInit {

  constructor(public router: Router, private api: ApiSerivice) { }
  public content: any = {
    legaltype:   '',
    legaltitle:  '',
    legalconent: ''
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
      this.router.navigate(['article'], {
        queryParams: {legaltype: this.content.legaltype}
        });
  }
}
