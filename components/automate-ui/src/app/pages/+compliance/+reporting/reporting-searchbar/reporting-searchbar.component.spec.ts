import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { TestBed, ComponentFixture } from '@angular/core/testing';
import { ReportingSearchbarComponent } from './reporting-searchbar.component';
import { using } from 'app/testing/spec-helpers';
import * as moment from 'moment';
import {
  ReportQueryService
} from 'app/pages/+compliance/shared/reporting';

describe('ReportingSearchbarComponent', () => {
  let fixture: ComponentFixture<ReportingSearchbarComponent>,
    component: ReportingSearchbarComponent;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [
      ],
      declarations: [
        ReportingSearchbarComponent
      ],
      providers: [
        ReportQueryService
      ],
      schemas: [ CUSTOM_ELEMENTS_SCHEMA ]
    });

    fixture = TestBed.createComponent(ReportingSearchbarComponent);
    component = fixture.componentInstance;
  });


  it('initializes', () => {
    expect(component).not.toBeNull();
  });


  describe('onDaySelect()', () => {
    using([
      ['beginning of the month', '2019-10-01', 9],
      ['currently selected day', '2019-10-23', 9],
      ['day in month', '2019-10-28', 9],
      ['end of the month', '2019-10-31', 9],
      ['end of the previous month', '2019-09-30', 8],
      ['beginning of the next month', '2019-11-01', 10]
    ], function (description: string, dateString: string, month) {
      it('selection of ' + description, () => {
        component.visibleDate = moment('20191023', 'YYYYMMDD');
        component.onDaySelect(dateString);

        expect(component.visibleDate.month()).toEqual(month);
      });
    });
  });
});
