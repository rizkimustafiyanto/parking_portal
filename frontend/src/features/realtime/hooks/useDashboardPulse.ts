"use client"

import { useEffect, useRef, useState } from "react"

import type { DashboardRole } from "../types"

type DashboardPulseState = {
  snapshot: {
    role: DashboardRole
    updatedAt: number | null
    pulseCount: number
    invoiceCount: number
    paymentCount: number
  }
  isRefreshing: boolean
  isPulsing: boolean
  pulseCount: number
  lastPulseAt: number | null
  triggerPulse: () => void
}

function useDashboardPulse(role: DashboardRole): DashboardPulseState {
  const [isPulsing, setIsPulsing] = useState(false)
  const [pulseCount, setPulseCount] = useState(0)
  const [lastPulseAt, setLastPulseAt] = useState<number | null>(null)
  const intervalRef = useRef<number | null>(null)

  const clearPulseInterval = () => {
    if (intervalRef.current !== null) {
      window.clearInterval(intervalRef.current)
      intervalRef.current = null
    }
  }

  const triggerPulse = () => {
    setIsPulsing(true)
    setPulseCount((current) => current + 1)
    setLastPulseAt(Date.now())

    window.setTimeout(() => {
      setIsPulsing(false)
    }, 350)
  }

  useEffect(() => {
    clearPulseInterval()
    intervalRef.current = window.setInterval(triggerPulse, role === "member" ? 5000 : 6500)

    return clearPulseInterval
  }, [role])

  return {
    snapshot: {
      role,
      updatedAt: lastPulseAt,
      pulseCount,
      invoiceCount: 0,
      paymentCount: 0,
    },
    isRefreshing: isPulsing,
    isPulsing,
    pulseCount,
    lastPulseAt,
    triggerPulse,
  }
}

export { useDashboardPulse }
