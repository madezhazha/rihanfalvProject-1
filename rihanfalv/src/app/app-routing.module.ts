import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { FeedbackComponent } from './components/feedback/feedback.component';
import { FeedbacksuccessComponent } from './components/Feedback/feedbacksuccess/feedbacksuccess.component';
import { PersonalfeedbackComponent } from './components/Feedback/personalfeedback/personalfeedback.component';

const routes: Routes = [
  {
    path:'Feedback',component:FeedbackComponent
  },
  {
    path:'Personalfeedback',component:PersonalfeedbackComponent
  },
  {
    path:'Feedbacksuccess',component:FeedbacksuccessComponent
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
