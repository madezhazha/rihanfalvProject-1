import { Component, OnInit } from '@angular/core';

import { DataService } from '../../../services/data.service';
import { Router, NavigationExtras } from '@angular/router';
@Component({
  selector: 'app-popular',
  templateUrl: './popular.component.html',
  styleUrls: ['./popular.component.scss']
})
export class PopularComponent implements OnInit {

  private threadList: any;
  private isMax: boolean = false;
  private nowData: any;


  constructor(
    private dataService: DataService,
    private router: Router
  ) { }

  ngOnInit() {
    this.dataService.getThreadList().subscribe((response) => {
      this.threadList = response;
      console.log(this.threadList);
      this.nowData = this.threadList.slice(0,2);
      console.log(this.threadList);
      if (this.nowData.length == this.threadList.length) {
        this.isMax = true;
      } else {
        this.isMax = false;
      }
    });
  }

  read(item) {
    let navigationExtras: NavigationExtras = {
      queryParams: {
        topicID: item.thread.ID,
      },
    }
    this.router.navigate(['/post'], navigationExtras)
  }

  more() {
    this.nowData = this.threadList.slice(0, this.nowData.length + 5);
    console.log(this.nowData);
    if (this.nowData.length == this.threadList.length) {
      this.isMax = true;
    }
  }

}
