import React, { useEffect, useState, useRef } from 'react'


export function useFetch<T>(url: string): { data: T | null; error: string | null; loading: boolean } {
    const [data, setData] = useState<T | null>(null)
    const [loading, setLoading] = useState<boolean>(false)
    const [error, setError] = useState<string | null>(null);
    const abortControllerRef = useRef<AbortController | null>(null);

    useEffect(() => {
        const controller = new AbortController();
        abortControllerRef.current = controller;
        setLoading(true)
        fetch(url, { signal: controller.signal })
            .then((response) => {
                if (!response.ok) {
                    throw new Error(`Error: ${response.statusText}`);
                }
                return response.json();
            }).then((result) => {
                setData(result);
            }).catch((err) => {
                if (err.name !== 'AbortError') {
                    setError(err.message);
                }
            }).finally(() => {
                setLoading(false);
            });


        return () => controller.abort()
    }, [url])

    return {
        data,
        error,
        loading
    }
}