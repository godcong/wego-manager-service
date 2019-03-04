import {Component, OnChanges, OnInit, SimpleChanges} from '@angular/core';
import {HomeDataService} from './home-data.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit, OnChanges {
  private links: string[] = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '9', '9', '9', '9', '9', '9', '9', '9', '9', '9'];
  private data: HomeDataService;
  public activities: any;

  constructor(data: HomeDataService) {
    this.data = data;
    this.activities = this.data.getActivityList();
  }

  ngOnInit() {
    // this.activities = this.data.getActivityList();
  }


  ngOnChanges(changes: SimpleChanges): void {
    // this.activities = this.data.getActivityList();
  }

}
