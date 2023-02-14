import React from 'react';
import {RouteObject} from 'react-router-dom';

import { RiskAversion } from './RiskAversion';

export const routes: RouteObject = {
  path: '/experiments',
  children: [
    {
      path: 'risk-aversion',
      element: <RiskAversion />,
    }
  ],
};
