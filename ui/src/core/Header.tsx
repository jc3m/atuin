import React from 'react';
import { Link } from 'react-router-dom';

interface NavigationLinkProps extends React.PropsWithChildren {
  to: string;
}

const NavigationLink: React.FC<NavigationLinkProps> = ({ to, children }) => {
  return (
    <Link
      to={to}
      className="px-4 text-center text-zinc-600 font-medium hover:text-black transition"
    >
      {children}
    </Link>
  );
};

export const Header = () => {
  return (
    <header className="flex flex-col items-center justify-center py-2">
      <h1 className="text-4xl font-bold font-sans">Cracker Barrel Brain Trust</h1>
      <p className="font-serif italic text-zinc-500 text-sm my-0.5">Putting the "Ed" in "Education"</p>
      <nav className="flex items-center justify-center mt-2">
        <NavigationLink to="/">Home</NavigationLink>
        <NavigationLink to="/experiments">Experiments</NavigationLink>
        <NavigationLink to="#">Podcasts</NavigationLink>
      </nav>

      <div className="w-full border-b border-zinc-200 mt-1" />
    </header>
  );
};
