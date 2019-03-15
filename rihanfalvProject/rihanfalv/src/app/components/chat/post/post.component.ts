import { Component, OnInit } from '@angular/core';

import { ActivatedRoute } from '@angular/router';
import { DataService } from '../../../services/data.service';
import { Router, NavigationExtras } from '@angular/router';

import { InputData } from '../../head/langing/land/input'
import { DomSanitizer } from '@angular/platform-browser';

@Component({
  selector: 'app-post',
  templateUrl: './post.component.html',
  styleUrls: ['./post.component.css']
})
export class PostComponent implements OnInit {

  private

  // 楼主的id
  private hostData: any;

  private thread: any;
  private post: any;
  private isMax: boolean = false;
  private nowData: any;

  //  是否收藏
  private collection: any;

  // 是否登录
  private Islogin: boolean;
  // 是否弹出登录框
  private IfWantLogin: boolean = false;

  // 登录的用户ID
  private userID: any;

  constructor(
    private activatedRouter: ActivatedRoute,
    private dataService: DataService,
    private router: Router,
    private sanitizer: DomSanitizer,
  ) { }

  ngOnInit() {
    this.activatedRouter.queryParams.subscribe((response) => {
      // localStorage.removeItem("id");

      //获取登陆用户信息
      this.userID = localStorage.getItem("id");
      if (this.userID != null) {
        this.Islogin = true;
      } else {
        this.Islogin = false
        this.userID = -1;
      }
      // debugger;
      this.hostData = response;

      // 这里的userID登陆的用户ID(传回后台判断是否收藏了此贴)
      this.dataService.getPostList(this.userID, response.topicID).subscribe((resp) => {
        this.collection = resp.collection;
        this.post = resp.post;
        this.thread = resp.thread;

        if (this.thread.user.Image.indexOf("assets") == -1) {
          let temp = 'data:image/png;base64, ' + this.thread.user.Image; //给base64添加头缀
          this.thread.user.Image = this.sanitizer.bypassSecurityTrustUrl(temp);
        }

        for (let i = 0; i < this.post.length; i++) {
          const element = this.post[i];
          if (element.user.Image.indexOf("assets") == -1) {
            let temp = 'data:image/png;base64, ' + element.user.Image; //给base64添加头缀
            element.user.Image = this.sanitizer.bypassSecurityTrustUrl(temp);
          }
        }

        if (this.post) {
          this.nowData = this.post.slice(0, 4);
          if (this.nowData.length == this.post.length) {
            this.isMax = true;
          } else {
            this.isMax = false;
          }
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
    if (this.post) {
      this.nowData = this.post.slice(0, this.nowData.length + 5);
      console.log(this.post);
      if (this.nowData.length == this.post.length) {
        this.isMax = true;
      }
    } else {
      this.isMax = true;
    }
  }

  collect() {
    // 1也是模拟用户id
    this.dataService.collect(this.userID, this.hostData.topicID).subscribe((response) => {
      this.collection.Collectioncontentid = response;
    });
  }

  cancel() {
    this.dataService.cancel(this.userID, this.hostData.topicID).subscribe((response) => {
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
