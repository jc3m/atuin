import React from 'react';
import {Link} from 'react-router-dom';

export const Experiments: React.FC = () => {
  return (
    <section>
      <h1 className="text-3xl">Experiments</h1>
      <ul className="list-disc list-inside">
        <li className="underline">
          <Link to="risk-aversion">Risk Aversion</Link>
        </li>
      </ul>
    </section>
  );
}
