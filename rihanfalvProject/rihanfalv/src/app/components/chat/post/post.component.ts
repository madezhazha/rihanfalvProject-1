import { Component, OnInit } from '@angular/core';

import { ActivatedRoute } from '@angular/router';
import { DataService } from '../../../services/data.service';
import { Router, NavigationExtras } from '@angular/router';

@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.scss']
})
export class PostComponent implements OnInit {

  // 楼主的数据
  private hostData: any;

  private thread: any;
  private post: any;
  private isMax: boolean = false;

  private nowData: any;

  constructor(
    private activatedRouter: ActivatedRoute,
    private dataService: DataService,
    private router: Router,
  ) { }

  ngOnInit() {
    this.activatedRouter.queryParams.subscribe((response) => {
      // debugger;
      this.hostData = response;
      console.log(this.hostData.topicID);
      this.dataService.getPostList(response.topicID).subscribe((resp) => {
        this.post = resp.post;
        this.thread = resp.thread;
        this.nowData = this.post.slice(0, 4);
        console.log(resp);
        if (this.nowData.length == this.post.length) {
          this.isMax = true;
        }else{
          this.isMax = false;
        }
      });
    });
  }

  reply() {
    // 求出当前帖子的最大楼层数
    let maxFloor: number = 0;
    if (this.post != null) {
      this.post.forEach(element => {
        if (element.post.Floor > maxFloor) {
          maxFloor = element.post.Floor;
        }
      });
    } else {
      maxFloor = 1;
    }

    let navigationExtras: NavigationExtras = {
      queryParams: {
        topicID: this.hostData.topicID,
        topicTitle: this.thread.thread.Topictitle,
        floor: maxFloor + 1,  // 回答的帖子加载在页面的楼层数
      }
    }
    this.router.navigate(['/replyPage'], navigationExtras)
  }

  more() {
    this.nowData = this.post.slice(0, this.nowData.length + 5);
    console.log(this.post);
    if (this.nowData.length == this.post.length) {
      this.isMax = true;
    }
  }

}
