import { Component, OnInit } from '@angular/core';
import { Router} from '@angular/router';

@Component({
  selector: 'app-feedbacksuccess',
  templateUrl: './feedbacksuccess.component.html',
  styleUrls: ['./feedbacksuccess.component.css']
})
export class FeedbacksuccessComponent implements OnInit {

  constructor(public router:Router) { }

  ngOnInit() {
  }
  Historicalfeedback(){
    this.router.navigate(['/Personalfeedback/'])

  }
  jixu(){
    this.router.navigate(['/Feedback/'])
  }

}
