import React from 'react';
import {RouteObject} from 'react-router-dom';

import { Experiments } from '.';
import { RiskAversion } from './RiskAversion';

export const routes: RouteObject = {
  path: '/experiments',
  children: [
    {
      path: '',
      element: <Experiments />,
    },
    {
      path: 'risk-aversion',
      element: <RiskAversion />,
    }
  ],
};
