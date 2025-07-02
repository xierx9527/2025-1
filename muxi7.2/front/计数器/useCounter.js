import { useState, useEffect, useCallback } from 'react';

export function useCounter(initialValue = 0) {
    const [count, setCount] = useState(initialValue);

    const increment = useCallback(() => {
        setCount(prev => prev + 1);
    }, []);

    const decrement = useCallback(() => {
        setCount(prev => prev - 1);
    }, []);

    const reset = useCallback(() => {
        setCount(initialValue);
    }, [initialValue]);

    useEffect(() => {
        console.log(`计数器更新: ${count}`);
    }, [count]);

    return { count, increment, decrement, reset };
}