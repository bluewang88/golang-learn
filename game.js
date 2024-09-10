const ctx = wx.createCanvasContext('gameCanvas')

const bird = {
  x: 50,
  y: 150,
  velocity: 0,
  gravity: 0.5,
  jump: -10,
  radius: 20
}

const pipes = []

function gameLoop() {
  // 更新鸟的位置
  bird.velocity += bird.gravity
  bird.y += bird.velocity

  // 绘制背景
  ctx.fillStyle = '#70c5ce'
  ctx.fillRect(0, 0, 300, 512)

  // 绘制鸟
  ctx.fillStyle = '#yellow'
  ctx.beginPath()
  ctx.arc(bird.x, bird.y, bird.radius, 0, 2 * Math.PI)
  ctx.fill()

  // 绘制管道
  ctx.fillStyle = 'green'
  pipes.forEach(pipe => {
    ctx.fillRect(pipe.x, 0, pipe.width, pipe.top)
    ctx.fillRect(pipe.x, pipe.bottom, pipe.width, 512 - pipe.bottom)
    pipe.x -= 2
  })

  // 碰撞检测和游戏结束逻辑这里省略...

  ctx.draw()
  requestAnimationFrame(gameLoop)
}

// 点击屏幕时鸟跳跃
wx.onTouchStart(() => {
  bird.velocity = bird.jump
})

// 开始游戏循环
gameLoop()