import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { FeedbackComponent } from './components/feedback/feedback.component';
import { FeedbacksuccessComponent } from './components/Feedback/feedbacksuccess/feedbacksuccess.component';
import { PersonalfeedbackComponent } from './components/Feedback/personalfeedback/personalfeedback.component';
import { CaseComponent } from './components/case/case.component';
import { CaseDataComponent } from './components/case/case-data/case-data.component';

const routes: Routes = [
  // {
  //   path:"",redirectTo:"homedata",pathMatch:"full"
  // },  这个表示一开始显示的组件
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
    path:"CaseAnalysis",component:CaseComponent
  },
  {
    path:"display-data",component:CaseDataComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
