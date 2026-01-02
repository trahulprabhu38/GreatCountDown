import React from 'react';

const CountdownDisplay = ({ data, filter }) => {
  if (!data) return null;

  const getTitle = () => {
    switch (filter) {
      case 'year':
        return `Days left to complete ${data.year || '2026'}`;
      case 'month':
        return `Days left in ${data.month || 'this month'}`;
      case 'week':
        return `Days left this week`;
      default:
        return 'Days remaining';
    }
  };

  return (
    <div className="countdown-display">
      <h2 className="title">{getTitle()}</h2>
      <div className="countdown-number">{data.daysRemaining}</div>
      <div className="subtitle">
        {filter === 'year' && `out of ${data.daysTotal} days`}
        {filter === 'month' && `out of ${data.daysTotal} days in ${data.month}`}
        {filter === 'week' && `out of ${data.daysTotal} days this week`}
      </div>
    </div>
  );
};

export default CountdownDisplay;
