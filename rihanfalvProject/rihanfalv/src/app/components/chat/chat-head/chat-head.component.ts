import { Component, OnInit } from '@angular/core';

import { Router } from '@angular/router';

@Component({
  selector: 'app-chat-head',
  templateUrl: './chat-head.component.html',
  styleUrls: ['./chat-head.component.scss']
})
export class ChatHeadComponent implements OnInit {

  constructor(
    private router:Router,
  ) { }

  ngOnInit() {
  }

  popular(){
    this.router.navigate(['/discussionarea']);
  }
  tag(){
    this.router.navigate(['/tag']);
    
  }
}
