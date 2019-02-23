import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { FeedbackComponent } from './components/feedback/feedback.component';
import { FeedbacksuccessComponent } from './components/Feedback/feedbacksuccess/feedbacksuccess.component';
import { PersonalfeedbackComponent } from './components/Feedback/personalfeedback/personalfeedback.component';
import { PaperComponent } from './components/paper/paper.component';
import { PaperwebComponent } from './components/paper/paperweb/paperweb.component';

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
  {
    path:'paper',component:PaperComponent
  },
  {
    path:'paperweb/:ID',component:PaperwebComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
