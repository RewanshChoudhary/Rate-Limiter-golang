local capacity = tonumber(ARGV[1])
local fill_rate = tonumber(ARGV[2])
local now = tonumber(ARGV[3])
local requested = tonumber(ARGV[4])
local requested=tonumber(ARGV[5])

local last_filled=tonumber(redis.call('hget',key[1],'last_filled') or now )
local current_tokens=tonumber(redis.call('hget',key[1],'current_tokens') or capacity)

local elapsed=now-last_filled
local refill=elapsed*fillrate
current_tokens=math.min(refill+current_tokens,capacity)

if current_tokens>=requested then
    current_tokens=current_tokens-requested
    redis.call('hset', KEYS[1], 'current_tokens', current_tokens)
    redis.call('hset', KEYS[1], 'last_filled', now)
    return 1
     --Request Accepted    
else 
    redis.call('hset',KEYS[1],'current_tokens',current_tokens)
    redis.call('hset',KEYS[2],'last_filled',now)
    return 0
    --Request Denied

    