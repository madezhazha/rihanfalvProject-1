import { Injectable } from '@angular/core';
import { CanActivate  } from '@angular/router';
@Injectable({
  providedIn: 'root'
})
export class LoginserviceService implements CanActivate {
  constructor() {
  }
  canActivate() {
    if (localStorage.getItem('id')) {
      return true;
    } else {
      return false;
    }
    // return true;
  }
}
