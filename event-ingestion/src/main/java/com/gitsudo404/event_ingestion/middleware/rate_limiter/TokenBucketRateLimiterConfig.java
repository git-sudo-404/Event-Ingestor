/*
 * MIT License
 *
 * Copyright (c) 2026 git-sudo-404
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * Of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * Copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * Copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING, BUT NOT LIMITED TO, THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES, OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package com.gitsudo404.event_ingestion.middleware.rate_limiter;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.redis.core.StringRedisTemplate;

import com.gitsudo404.event_ingestion.middleware.rate_limiter.properties.VehicleBatteryEventTokenBucketProperties;
import com.gitsudo404.event_ingestion.middleware.rate_limiter.properties.VehicleDiagnosticEventTokenBucketProperties;
import com.gitsudo404.event_ingestion.middleware.rate_limiter.properties.VehicleEngineTelemetryEventTokenBucketProperties;
import com.gitsudo404.event_ingestion.middleware.rate_limiter.properties.VehicleFuelEventTokenBucketProperties;
import com.gitsudo404.event_ingestion.middleware.rate_limiter.properties.VehicleGPSEventTokenBucketProperties;

@Configuration
public class TokenBucketRateLimiterConfig {

  @Bean("vehicleGPSEventRateLimiter")
  RateLimiter vehicleGPSEventRateLimiter(VehicleGPSEventTokenBucketProperties properties,
      StringRedisTemplate redis) {
    return new TokenBucketRateLimiter(
        properties,
        redis,
        "rate-limit:token-bucket:vehicle:gps");
  }

  @Bean("vehicleBatteryEventRateLimiter")
  RateLimiter vehicleBatteryEventRateLimiter(
      VehicleBatteryEventTokenBucketProperties properties,
      StringRedisTemplate redis) {
    return new TokenBucketRateLimiter(properties, redis, "rate-limit:token-bucket:vehicle:battery");
  }

  @Bean("vehicleFuelEventRateLimiter")
  RateLimiter vehicleFuelEventRateLimiter(
      VehicleFuelEventTokenBucketProperties properties,
      StringRedisTemplate redis) {
    return new TokenBucketRateLimiter(properties, redis, "rate-limit:token-bucket:vehicle:fuel");
  }

  @Bean("vehicleDiagnosticEventRateLimiter")
  RateLimiter vehicleDiagnosticEventRateLimiter(
      VehicleDiagnosticEventTokenBucketProperties properties,
      StringRedisTemplate redis) {
    return new TokenBucketRateLimiter(properties, redis, "rate-limit:token-bucket:vehicle:diagnostic");
  }

  @Bean("vehicleEngineTelemetryEventRateLimiter")
  RateLimiter vehicleEngineTelemetryRateLimiter(
      VehicleEngineTelemetryEventTokenBucketProperties properties,
      StringRedisTemplate redis) {
    return new TokenBucketRateLimiter(properties, redis, "rate-limit:token-bucket:vehicle:engine-telemetry");
  }

}
