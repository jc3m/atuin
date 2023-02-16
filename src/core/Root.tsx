import React from 'react';
import { Outlet } from 'react-router';

import { Header } from './Header';

export const Root = () => {
  return (
    <div className="container mx-auto">
      <Header />
      <main className="my-2 px-4">
        <Outlet />
      </main>
    </div>
  )
};
