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

import java.util.List;
import java.util.Map;

import org.springframework.data.redis.core.StringRedisTemplate;

import com.gitsudo404.event_ingestion.middleware.rate_limiter.properties.TokenBucketRateLimiterProperties;

/*
 * We can calculate the currenlty available tokens in a bucket without physically refilling it, using the below formulae
 *
 * C : Maximum Capacity of the Bucket 
 * r : refill rate 
 * T : currentlt available tokens 
 *
 * T' = min( C , T + r * time_elapsed )
 * where, 
 *    time_elapsed is the time since the last refill calculation 
 */

public class TokenBucketRateLimiter implements RateLimiter {

  private final TokenBucketRateLimiterProperties properties;

  private final StringRedisTemplate stringRedisTemplate;

  private final String keyPrefix;

  public TokenBucketRateLimiter(TokenBucketRateLimiterProperties properties, StringRedisTemplate stringRedisTemplate,
      String keyPrefix) {
    this.properties = properties;
    this.stringRedisTemplate = stringRedisTemplate;
    this.keyPrefix = keyPrefix;
  }

  public Boolean allowRequest(String entityId) {

    String key = this.keyPrefix + ":" + entityId;
    List<Object> redisVals = stringRedisTemplate.opsForHash().multiGet(key, List.of("tokens", "lastRefillMs"));

    double currentMs = System.currentTimeMillis();

    double tokens = redisVals.get(0) == null ? properties.getCapacity()
        : Double.parseDouble(redisVals.get(0).toString());
    double lastRefillMs = redisVals.get(1) == null ? currentMs : Double.parseDouble(redisVals.get(1).toString());

    double elapsedSeconds = Math.max(
        0,
        currentMs - lastRefillMs) / 1000.0;

    double elapsedPeriods = elapsedSeconds / properties.getRefillPeriodS();

    double newTokens = Math.min(
        properties.getCapacity(),
        tokens + properties.getRefillRate() * elapsedPeriods);

    if (newTokens >= 1.0)
      newTokens -= 1.0;

    stringRedisTemplate.opsForHash().putAll(key, Map.of(
        "tokens", newTokens,
        "lastRefillMs", currentMs));

    return newTokens > 1.0;

  }

}
