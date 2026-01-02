import React from 'react';

const FilterTabs = ({ activeFilter, onFilterChange }) => {
  return (
    <div className="filter-tabs">
      <button
        className={activeFilter === 'year' ? 'active' : ''}
        onClick={() => onFilterChange('year')}
      >
        YEAR
      </button>
      <button
        className={activeFilter === 'month' ? 'active' : ''}
        onClick={() => onFilterChange('month')}
      >
        MONTH
      </button>
      <button
        className={activeFilter === 'week' ? 'active' : ''}
        onClick={() => onFilterChange('week')}
      >
        WEEK
      </button>
    </div>
  );
};

export default FilterTabs;
