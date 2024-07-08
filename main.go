package main

import (
    rl "github.com/gen2brain/raylib-go/raylib"
)

const maxColumns = 20

func main() {
    screenWidth := int32(800)
    screenHeight := int32(450)

    rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - 3d camera first person")
    defer rl.CloseWindow()

    camera := rl.Camera{}
    camera.Position = rl.NewVector3(0.0, 2.0, 4.0) // Camera position
    camera.Target = rl.NewVector3(0.0, 2.0, 0.0)   // Camera looking at point
    camera.Up = rl.NewVector3(0.0, 1.0, 0.0)       // Camera up vector (rotation towards target)
    camera.Fovy = 60.0                             // Camera field-of-view Y
    camera.Projection = rl.CameraPerspective             // Camera projection type
      // Initialize audio device
      rl.InitAudioDevice()
      defer rl.CloseAudioDevice()
  
      // Load explosion sound
      fxBoom := rl.LoadSound("resources/boom.wav")
      defer rl.UnloadSound(fxBoom)
  
      // Load explosion texture
      explosion := rl.LoadTexture("resources/explosion.png")
      defer rl.UnloadTexture(explosion)
  
      // Init variables for animation
      const NUM_FRAMES_PER_LINE = 5 // Example value, replace with actual number
      const NUM_LINES = 5           // Example value, replace with actual number
  
      frameWidth := float32(explosion.Width) / float32(NUM_FRAMES_PER_LINE)   // Sprite one frame rectangle width
      frameHeight := float32(explosion.Height) / float32(NUM_LINES)           // Sprite one frame rectangle height
      currentFrame := 0
      currentLine := 0

       // Define frame rectangle and position
    frameRec := rl.NewRectangle(0, 0, frameWidth, frameHeight)
    position := rl.NewVector2(0.0, 0.0)

    // Animation control variables
    active := false
    framesCounter := 0
  

    var heights [maxColumns]float32
    var positions [maxColumns]rl.Vector3
    var colors [maxColumns]rl.Color

    for i := 0; i < maxColumns; i++ {
        heights[i] = float32(rl.GetRandomValue(1, 12))
        positions[i] = rl.NewVector3(float32(rl.GetRandomValue(-15, 15)), heights[i]/2.0, float32(rl.GetRandomValue(-15, 15)))
        colors[i] = rl.NewColor(uint8(rl.GetRandomValue(20, 255)), uint8(rl.GetRandomValue(10, 55)), 30, 255)
    }

    model := rl.LoadModel("resources/plane.obj");                  // Load model
    texture := rl.LoadTexture("resources/plane_diffuse.png");  // Load model texture
    mat1 := model.GetMaterials()
    mat1[0].GetMap(rl.MapDiffuse).Texture = texture // Set model diffuse texture

    rl.DisableCursor() // Limit cursor to relative movement inside the window

    rl.SetTargetFPS(60) // Set our game to run at 60 frames-per-second

    // Main game loop
    for !rl.WindowShouldClose() {
        rl.UpdateCamera(&camera, rl.CameraFirstPerson) // Update camera
            // Check for mouse button pressed and activate explosion (if not active)
            if rl.IsKeyPressed(rl.KeySpace) && !active {
                active = true
    
                 position.X = float32(screenWidth) / 2.0
                 position.Y = float32(screenHeight) / 2.0
                 position.X -= frameWidth / 2.0
                 position.Y -= frameHeight / 2.0 - 90.0
    
                rl.PlaySound(fxBoom)
            }
    
            // Compute explosion animation frames
            if active {
                framesCounter++
    
                if framesCounter > 2 {
                    currentFrame++
    
                    if currentFrame >= NUM_FRAMES_PER_LINE {
                        currentFrame = 0
                        currentLine++
    
                        if currentLine >= NUM_LINES {
                            currentLine = 0
                            active = false
                        }
                    }
    
                    framesCounter = 0
                }
            }
    
            frameRec.X = frameWidth * float32(currentFrame)
            frameRec.Y = frameHeight * float32(currentLine)

        rl.BeginDrawing()

        rl.ClearBackground(rl.White)

        rl.BeginMode3D(camera)
		rl.DrawPlane(rl.NewVector3(0.0, 0.0, 0.0), rl.NewVector2(102.0, 102.0), rl.Brown)
        for i := 0; i < maxColumns; i++ {
            rl.DrawCube(positions[i], 2.0, heights[i], 2.0, colors[i])
            rl.DrawCubeWires(positions[i], 2.0, heights[i], 2.0, rl.Gray)
        }
        rl.DrawModel(model,camera.Target,0.03,rl.White)

        
        rl.EndMode3D()
        if active {
              rl.DrawTextureRec(explosion, frameRec, position, rl.White)
        }

        rl.EndDrawing()
    }
}