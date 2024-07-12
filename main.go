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
      gun := rl.LoadTexture("resources/shotgun.png")
      defer rl.UnloadTexture(gun)
      guyModel := rl.LoadModel("resources/guy.iqm")
      defer rl.UnloadModel(guyModel)
      guyTexture := rl.LoadTexture("resources/guytex.png")
      defer rl.UnloadTexture(guyTexture)
  
      // Set model material texture
      materials := guyModel.GetMaterials()
      materials[0].GetMap(rl.MapDiffuse).Texture = guyTexture
      rl.SetMaterialTexture(guyModel.Materials, rl.MapDiffuse, guyTexture)

         // Load animation data
    anims := rl.LoadModelAnimations("resources/guyanim.iqm")
    defer rl.UnloadModelAnimations(anims)
    var animFrameCounter int32 = 0

  
      // Init variables for animation
      const NUM_FRAMES_PER_LINE = 5 // Example value, replace with actual number
      const NUM_LINES = 1           // Example value, replace with actual number
  
    currentFrame := 0; 
    active:= false;

    framesCounter := 0;
    framesSpeed := 8;  
    frameRec := rl.NewRectangle(0.0, 0.0, float32(gun.Width)/NUM_FRAMES_PER_LINE, float32(gun.Height)/NUM_LINES)

    var positions [maxColumns]rl.Vector3

    for i := 0; i < maxColumns; i++ {

        positions[i] = rl.NewVector3(float32(rl.GetRandomValue(-50, 50)), -0.5, float32(rl.GetRandomValue(-50, 50)))

    }

    // guyModel := rl.LoadModel("resources/plane.obj");                  // Load model
    // guyTexture := rl.LoadTexture("resources/plane_diffuse.png");  // Load model texture
    // mat1 := guyModel.GetMaterials()
    // mat1[0].GetMap(rl.MapDiffuse).Texture = guyTexture // Set model diffuse texture

    rl.DisableCursor() // Limit cursor to relative movement inside the window

    rl.SetTargetFPS(60) // Set our game to run at 60 frames-per-second

    // Main game loop
    for !rl.WindowShouldClose() {
        rl.UpdateCamera(&camera, rl.CameraFirstPerson) // Update camera

     //==============================================================================
     // BEGIN GUN ANIMATION LOGIC  
            // Check for mouse button pressed and activate explosion (if not active)
             if rl.IsMouseButtonDown(rl.MouseButtonLeft) && !active {
                rl.PlaySound(fxBoom)
                active = true;}
     if active {
                 // Update frames counter and animation frame
        framesCounter++
        if framesCounter >= (60 / framesSpeed) {
            framesCounter = 0
            currentFrame++

            if currentFrame >= NUM_FRAMES_PER_LINE {
                currentFrame = 0
                active = false
            }

            frameRec.X = float32(currentFrame) * frameRec.Width
        }
    
    }
    // END GUN ANIMATION LOGIC
//==============================================================================
//==============================================================================
// BEGIN ZOMBIE GUY ANIMATION LOGIC     
    animFrameCounter++
        rl.UpdateModelAnimation(guyModel, anims[0], animFrameCounter)
        if animFrameCounter >= anims[0].FrameCount {
            animFrameCounter = 0
        }

 // END ZOMBIE GUY  ANIMATION LOGIC
 //==============================================================================
        rl.BeginDrawing()

        rl.ClearBackground(rl.Black)

        rl.BeginMode3D(camera)
		rl.DrawPlane(rl.NewVector3(0.0, 0.0, 0.0), rl.NewVector2(102.0, 102.0), rl.Brown)
        // Draw the model with rotation and scaling
        rotationAxis := rl.NewVector3(1.0, 0.0, 0.0)
        scale := rl.NewVector3(1.0, 1.0, 1.0)
        for i := 0; i < maxColumns; i++ {
            rl.DrawModelEx(guyModel, positions[i], rotationAxis, -90.0, scale, rl.White)
        }

        
        rl.EndMode3D()
        
              rl.DrawTextureRec(gun, frameRec, rl.NewVector2(float32(screenWidth/2),float32(screenHeight/2)), rl.White)
       
         

        rl.EndDrawing()
    }
} 