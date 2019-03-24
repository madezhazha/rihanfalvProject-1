import { Component, OnInit } from '@angular/core';

import { DataService } from '../../../services/data.service';

import { Router, NavigationExtras } from '@angular/router';
import { ActivatedRoute } from '@angular/router';
@Component({
  selector: 'app-reply-page',
  templateUrl: './reply-page.component.html',
  styleUrls: ['./reply-page.component.css']
})
export class ReplyPageComponent implements OnInit {

  public RoutingData: any;
  // 回答的文本
  public text: string;

  // 登录的用户ID
  public userID: any;

  constructor(
    public activatedRouter: ActivatedRoute,
    public dataService: DataService,
    public router: Router,
  ) { }

  ngOnInit() {
    this.activatedRouter.queryParams.subscribe((response) => {
      this.userID=localStorage.getItem("id")
      console.log(this.userID);
      this.RoutingData = response;
    });
  }

  submit() {
    if (this.text == undefined || this.text == "") {
      alert("内容不能为空");
      return;
    }
    // 这里的userID代表回答的用户
    this.dataService.submitReply(this.userID, this.RoutingData.topicID,
      this.text, this.RoutingData.floor).subscribe(() => {
        let navigationExtras: NavigationExtras = {
          queryParams: {
            topicID: this.RoutingData.topicID,
          }
        }
        this.router.navigate(['/post'], navigationExtras)
      });

  }



}
