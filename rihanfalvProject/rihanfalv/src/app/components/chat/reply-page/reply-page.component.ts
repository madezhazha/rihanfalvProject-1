import { Component, OnInit } from '@angular/core';

import { DataService } from '../../../services/data.service';

import { Router, NavigationExtras } from '@angular/router';
import { ActivatedRoute } from '@angular/router';
@Component({
  selector: 'app-reply-page',
  templateUrl: './reply-page.component.html',
  styleUrls: ['./reply-page.component.scss']
})
export class ReplyPageComponent implements OnInit {

  private RoutingData: any;
  private text: string;

  constructor(
    private activatedRouter: ActivatedRoute,
    private dataService: DataService,
    private router: Router,
  ) { }

  ngOnInit() {
    this.activatedRouter.queryParams.subscribe((response) => {
      this.RoutingData = response;
      console.log(this.RoutingData);
    });
  }

  submit() {
    if (this.text == undefined || this.text == "") {
      alert("内容不能为空");
      return;
    }
    // 这里的1是模拟用户的userid
    this.dataService.submitReply("1", this.RoutingData.topicID,
      this.text, this.RoutingData.floor).subscribe(()=>{
        let navigationExtras: NavigationExtras = {
          queryParams: {
            topicID: this.RoutingData.topicID,
          }
        }
        this.router.navigate(['/post'], navigationExtras)
      });

  }



}
