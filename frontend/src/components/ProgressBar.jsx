import React from 'react';

const ProgressBar = ({ percentComplete }) => {
  const percentage = percentComplete || 0;
  const isSmallPercentage = percentage < 5;

  return (
    <div className="progress-container">
      <div className="progress-bar">
        <div
          className="progress-fill"
          style={{ width: `${percentage}%` }}
        >
          {!isSmallPercentage && (
            <span className="progress-text">{percentage.toFixed(2)}%</span>
          )}
        </div>
        {isSmallPercentage && (
          <span className="progress-text-outside">{percentage.toFixed(2)}%</span>
        )}
      </div>
      <div className="progress-label">Progress</div>
    </div>
  );
};

export default ProgressBar;
