import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {WebTokenService} from './web-token.service';

const HOST = 'http://localhost:8081';

interface UserActivity {
  PropertyID: string;
  ActivityID: string;
  UserID: string;
  IsStar: boolean;
  SpreadCode: string;
  IsPass: boolean;
  SpreadNumber: number;
}

interface UserActivityInterface {
  Current: number;
  Desc: boolean;
  Detail: UserActivity;
  Limit: number;
  Total: number;
  TotalPage: number;
}

@Injectable({
  providedIn: 'root'
})
export class DataService {
  private client: HttpClient;

  constructor(client: HttpClient) {
    this.client = client;
  }

  getActivityList() {
    return this.client.get(HOST + '/api/v0/spread/activity', {
        headers: {
          'Content-Type': 'application/json',
          token: WebTokenService.getToken(),
        }
      }
    );
  }

  getActivityInfo(id: string) {
    return this.client.get(HOST + '/api/v0/spread/activity/' + id, {
        headers: {
          'Content-Type': 'application/json',
          token: WebTokenService.getToken(),
        }
      }
    );
  }

  getUserActivityList() {
    return this.client.get(HOST + '/api/v0/spread/user/activity', {
        headers: {
          'Content-Type': 'application/json',
          token: WebTokenService.getToken(),
        }
      }
    );
  }

  getSpreadShareInfo(id: string, user: string) {
    return this.client.get(HOST + '/api/v0/spread/spread/' + id + '/share', {
        params: {
          user,
        },
        headers: {
          'Content-Type': 'application/json',
          token: WebTokenService.getToken(),
        }
      }
    );
  }

  getMyInfo() {
    return this.client.get(HOST + '/api/v0/spread/user/info', {
        headers: {
          'Content-Type': 'application/json',
          token: WebTokenService.getToken(),
        }
      }
    );
  }

  getMyActivity() {
    return this.client.get(HOST + '/api/v0/spread/user/activity', {
        headers: {
          'Content-Type': 'application/json',
          token: WebTokenService.getToken(),
        }
      }
    );
  }


  getMySpread() {
    return this.client.get(HOST + '/api/v0/spread/user/spread', {
        headers: {
          'Content-Type': 'application/json',
          token: WebTokenService.getToken(),
        }
      }
    );
  }

}
