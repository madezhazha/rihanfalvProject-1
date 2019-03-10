import { Component, OnInit } from '@angular/core';

import { ActivatedRoute } from '@angular/router';
import { DataService } from '../../../services/data.service';
import { Router, NavigationExtras } from '@angular/router';

import { InputData } from '../../head/langing/land/input'

@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.css']
})
export class PostComponent implements OnInit {

  private

  // 楼主的数据
  private hostData: any;

  private thread: any;
  private post: any;
  private isMax: boolean = false;
  private nowData: any;


  //  是否收藏
  private collection: any;

  // 是否登录
  private Islogin: boolean;
  private IfWantLogin: boolean = false;

  constructor(
    private activatedRouter: ActivatedRoute,
    private dataService: DataService,
    private router: Router,
  ) { }

  ngOnInit() {
    this.activatedRouter.queryParams.subscribe((response) => {
      // localStorage.removeItem("id"); 
      //获取用户信息
      let userid: string = localStorage.getItem("id");
      if (userid != null) {
        this.Islogin = true;
      } else {
        this.Islogin = false
      }
      // debugger;
      this.hostData = response;
      // console.log(this.hostData.topicID);
      // 1是假数据，代表用户id
      this.dataService.getPostList(1, response.topicID).subscribe((resp) => {
        this.collection = resp.collection;
        this.post = resp.post;
        this.thread = resp.thread;
        if (this.post != null) {
          this.nowData = this.post.slice(0, 4);
          if (this.nowData.length == this.post.length) {
            this.isMax = true;
          } else {
            this.isMax = false;
          }
          // console.log("hello")
        } else {
          this.isMax = true;
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

  collect() {
    // 1也是模拟用户id
    this.dataService.collect("1", this.hostData.topicID).subscribe((response) => {
      this.collection.Collectioncontentid = response;
    });
  }

  cancel() {
    this.dataService.cancel("1", this.hostData.topicID).subscribe((response) => {
      this.collection.Collectioncontentid = response;
    });
  }

  //关闭登陆框
  boxClose() {
    this.IfWantLogin = false;
  }

  getLoginData(loginData: InputData) {
    if (loginData.IfLogin == true) {
      this.Islogin = true
      this.boxClose();
    }
  }

  popupLogin() {
    this.IfWantLogin = true;
  }

}
